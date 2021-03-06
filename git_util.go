package main

import (
  "io/ioutil"
  "strings"
  "github.com/libgit2/git2go"
)

type CommitDiff struct{
  CommitId *git.Oid
  DiffString string
  CommitedOn string
}

var repo *git.Repository

func fileOperation(filename string, operation string, idx *git.Index) error {
  var err error
  if operation == "add" {
    err = idx.AddByPath(filename)
  } else if operation == "remove" {
    err = idx.RemoveByPath(filename)
  }
  return err
}

func getTree(filename string, operation string) (*git.Tree, error) {
  idx, err := repo.Index(); if err != nil {
    return nil, err
  }
  if err = fileOperation(filename, operation, idx); err != nil {
    return nil, err
  }
  err = idx.Write(); if err != nil {
    return nil, err
  }
  treeId, err := idx.WriteTree(); if err != nil {
    return nil, err
  }
  tree, err := repo.LookupTree(treeId); if err != nil {
    return nil, err
  }
  return tree, nil
}

func createFirstCommit(filename string, operation string) error{
    err := ioutil.WriteFile(*repoPath + "/" + filename, nil, 0600); if err != nil {
      return err
    }
    signature, err := repo.DefaultSignature(); if err != nil {
      return err
    }
    tree, err := getTree(filename, operation); if err != nil {
      return err
    }
    _ , err = repo.CreateCommit("refs/heads/master", signature, signature, "First Commit", tree); if err != nil {
      return err
    }
    return nil
}

func CreateRepository() error {
  var err error
  repo, err = git.InitRepository(*repoPath, false); if err != nil {
    return err
  }
  if _, err = repo.Head(); err != nil {
    err = createFirstCommit(".gitignore", "add"); if err != nil {
      return err
    }
  }
  return  nil
}

func AddFileInRepo(filename string, operation string) error {
  signature, err := repo.DefaultSignature(); if err != nil {
    return err
  }
  tree, err := getTree(filename, operation); if err != nil {
    return err
  }
  head, err := repo.Head()
  if err != nil {
    return err
  }
  commitTarget, err := repo.LookupCommit(head.Branch().Target()); if err != nil {
    return err
  }
  _ , err = repo.CreateCommit("refs/heads/master", signature, signature, filename, tree, commitTarget); if err != nil {
    return err
  }
  return nil
}

func walk() (*git.RevWalk, error){
  repo_walk, err := repo.Walk(); if err != nil{
    return nil, err
  }
  head, err := repo.Head(); if err != nil{
    return nil, err
  }
  err = repo_walk.Push(head.Branch().Target()); if err != nil{
    return nil, err
  }
  return repo_walk, nil
}

func GetData(rawId string) ([]byte, error){
  commit_id, err := git.NewOid(rawId); if err != nil {
      return nil, err
  }
  data, err := findContentByCommitId(commit_id); if err != nil {
      return nil, err
  }
  return data, nil
}

func findContentByCommitId(commitId *git.Oid) ([]byte,error){
  commit, err := repo.LookupCommit(commitId); if err != nil {
    return nil, err
  }
  tree, err := commit.Tree(); if err != nil {
    return nil, err
  }
  treeByName := tree.EntryByName(commit.Message()); if treeByName == nil {
    return nil, nil
  }
  blob, err := repo.LookupBlob(treeByName.Id); if err != nil {
    return nil, err
  }
  return blob.Contents(), nil
}

func FindCommitsInFile(filename string) ([]*git.Commit,error) {
  repo_walk, err := walk(); if err != nil {
    return nil, err
  }
  oid := new(git.Oid)
  array := make([]*git.Commit, 0)
  for {
    err = repo_walk.Next(oid); if err != nil{
      return array,nil
    }
    commit,err := repo.LookupCommit(oid); if err != nil{
      return array,nil
    }
    if commit.Message() == filename {
      array = append(array,commit)
    }
  }
  return array, nil
}

func getDiff(curr,old *git.Commit) (*git.Diff, error) {
  diffOpt , _ := git.DefaultDiffOptions()
  var old_tree *git.Tree
  var err error
  if old != nil {
    old_tree, err = old.Tree(); if err != nil {
      return nil, err
    }
  }
  curr_tree, err := curr.Tree(); if err != nil {
    return nil,err
  }
  diff, err := repo.DiffTreeToTree(old_tree, curr_tree, &diffOpt); if err != nil {
    return nil, err
  }
  return diff, nil
}

func GetDiffInFile(curr,old *git.Commit,filename string) (string,error){
  diff, err := getDiff(curr, old); if err != nil || diff == nil {
    return "", err
  }
  count,err := diff.NumDeltas(); if err != nil {
    return "",err
  }
  filterString := "diff --git a/" + filename + " b/" + filename
  for i := 0 ; i < count ; i ++ {
    patch, err := diff.Patch(i); if err != nil {
      return "", err
    }
    patchStr,err := patch.String(); if err != nil {
      return "", err
    }
    if (strings.HasPrefix(patchStr,filterString)){
     return patchStr,nil
    }
  }
  return "",nil
}
