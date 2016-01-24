package main

import(
  "path"
  "io/ioutil"
  "github.com/libgit2/git2go"
  "errors"
  "strings"
  "os"
)

type FileInfo struct {
  FileName string
  ModTime string
}
type Files struct {
  List []FileInfo
}

func (p * Page) Save() error {
  filename := path.Join(*repoPath, p.Title)
  if _, err := os.Stat(filename); err == nil {
    reader,_ := ioutil.ReadFile(filename)
    if (string(reader) == string(p.Body)){
      return nil
    }
  }
  err := ioutil.WriteFile(filename, p.Body, 0600)
  if(err != nil){
    return err
  }
  return AddFileInRepo(p.Title, "add")
}

func GetPayload(title string) (*Page, error) {
  file := path.Join(*repoPath, title)
  body, err := ioutil.ReadFile(file); if err != nil {
        return nil,err
  }
  return &Page{Title: title, Body: body}, nil
}

func GetFileList(dir string) ([]FileInfo, error) {
    files, err := ioutil.ReadDir(path.Join(dir, "/")); if err != nil {
      return nil, err
    }
    var file_list []FileInfo
    for _, f := range files {
      if f.Name() != ".git" && f.Name() != ".gitignore" {
        file := FileInfo{FileName: f.Name(), ModTime: f.ModTime().UTC().Format("3:04pm on Jan 2, 2006 (MST)")}
        file_list = append(file_list, file)
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

func GetFileAndRawId(path string, apiName string) (string, string){
  index := strings.LastIndex(path, "/")
  filename := path[len(apiName):index]
  rawId := path[index+1:]
  return filename, rawId
}

func CheckFileName(file_list []FileInfo, filename string) bool {
  for _, file := range file_list {
      if file.FileName == filename {
        return false
      }
  }
  return true
}

func RemoveFile(filename string) error {
  filepath := path.Join(*repoPath, filename)
  if err := os.Remove(filepath); err != nil {
    return err
  }
  return AddFileInRepo(filename, "remove")
}
