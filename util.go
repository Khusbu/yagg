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
  return AddFileInRepo(p.Title)
}

func GetPayload(title string) (*Page, error) {
  file := path.Join(*repoPath, title)
  body, err := ioutil.ReadFile(file); if err != nil {
        return nil,err
  }
  return &Page{Title: title, Body: body}, nil
}

func GetFileList(dir string) ([]string, error) {
    files, err := ioutil.ReadDir(path.Join(dir, "/")); if err != nil {
      return nil, err
    }
    var file_list []string
    for _, f := range files {
      if f.Name() != ".git" && f.Name() != ".gitignore" {
        file_list = append(file_list, f.Name())
      }
    }
    return file_list, nil
}
