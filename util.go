package main

import(
  "path"
  "io/ioutil"
  "github.com/libgit2/git2go"
  "errors"
)

func (p * Page) Save() error {
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

func createDiffArray(array []*git.Commit) []CommitDiff {
   if len(array) < 2 {
     return nil
   }
   diffArray := make([]CommitDiff,0)
   length := len(array)
   var cd CommitDiff
   for i := 1 ; i< length;  i++ {
     cd.CommitId = array[i-1].Id()
     cd.CommitedOn = array[i-1].Committer().When.UTC().Format("3:04pm on Jan 2, 2006 (MST)")
     patchStr,_ := GetDiffInFile(array[i-1],array[i],array[i].Message())
     cd.DiffString = patchStr
     diffArray = append(diffArray,cd)
   }
   return diffArray
}

func GetFirstCommit(firstCommit *git.Commit, filename string) (CommitDiff, error) {
  diffString, err := GetDiffInFile(firstCommit, firstCommit.Parent(0), filename); if err != nil {
    return CommitDiff{}, err
  }
  return CommitDiff{CommitId: firstCommit.Id(), DiffString: diffString,CommitedOn: firstCommit.Committer().When.UTC().Format("3:04pm on Jan 2, 2006 (MST)")}, nil
}

func GetHistory(title string) (*History, error) {
  array, err := FindCommitsInFile(title); if err != nil {
        return nil,err
  }
  if len(array) == 0 {
    return nil, errors.New("Page not found")
  }
  cdArray := createDiffArray(array)
  firstCommitDiff, err := GetFirstCommit(array[len(array)-1], title); if err != nil {
    return nil, err
  }
  cdArray = append(cdArray, firstCommitDiff)
  return &History{Title: title, CDiffs: cdArray}, nil
}
