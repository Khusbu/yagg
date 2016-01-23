package main

import (
  "io/ioutil"
  "github.com/libgit2/git2go"
)

func getTree(repo *git.Repository, filename string) (*git.Tree, error) {
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

func createFirstCommit(filename string, repo *git.Repository) error{
    err := ioutil.WriteFile(*repoPath + "/" + filename, nil, 0600); if err != nil {
      return err
    }
    signature, err := repo.DefaultSignature(); if err != nil {
      return err
    }
    tree, err := getTree(repo, filename); if err != nil {
      return err
    }
    _ , err = repo.CreateCommit("refs/heads/master", signature, signature, "First Commit", tree); if err != nil {
      return err
    }
    return nil
}

func CreateRepository() (*git.Repository, error) {
  repo, err := git.InitRepository(*repoPath, false); if err != nil {
    return nil, err
  }
  if _, err = repo.Head(); err != nil {
    err = createFirstCommit(".gitignore",repo); if err != nil {
      return nil, err
    }
  }
  return repo, nil
}
