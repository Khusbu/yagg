package main

import(
  "path"
  "io/ioutil"
)

func (p * Page) save() error {
  filename := path.Join(*repoPath, p.Title)
  err := ioutil.WriteFile(filename, p.Body, 0600)
  if(err != nil){
    return err
  }
  return nil
}
