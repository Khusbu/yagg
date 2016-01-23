package main

import (
  "io/ioutil"
  "github.com/libgit2/git2go"
)

var repo *git.Repository

func getTree(filename string) (*git.Tree, error) {
  idx, err := repo.Index(); if err != nil {
    return nil, err
  }
  err = idx.AddByPath(filename); if err != nil {
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

func createFirstCommit(filename string) error{
    err := ioutil.WriteFile(*repoPath + "/" + filename, nil, 0600); if err != nil {
      return err
    }
    signature, err := repo.DefaultSignature(); if err != nil {
      return err
    }
    tree, err := getTree(filename); if err != nil {
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
    err = createFirstCommit(".gitignore"); if err != nil {
      return err
    }
  }
  return  nil
}

func AddFileInRepo(filename string) error {
  signature, err := repo.DefaultSignature(); if err != nil {
    return err
  }
  tree, err := getTree(filename); if err != nil {
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
