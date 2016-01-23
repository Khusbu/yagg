package main

import (
  "flag"
  "fmt"
  "net/http"
  "log"
  "github.com/libgit2/git2go"
)

var (
	port        = flag.String("p", "8080", "Port to run yagg")
	host        = flag.String("b", "localhost", "Hostname to be used")
  repoPath    = flag.String("r", "data", "Set the git repositories path where data will be saved")
)

func createRepository() *git.Repository {
  repo, err := git.InitRepository(*repoPath, false)
  if err != nil {
    panic(err)
  }
  return repo
}

func main() {
  flag.Parse()

  repo := createRepository()
  fmt.Println("Repository working directory set as:", repo.Workdir())

  addr := fmt.Sprintf("%s:%s", *host, *port)
  fmt.Println("Listening on ", addr)
  
  if err := http.ListenAndServe(addr, nil); err != nil{
    log.Fatal(err)
  }
}
