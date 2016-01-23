package main

import (
  "flag"
  "fmt"
  "net/http"
  "log"
)

var (
	port        = flag.String("p", "8080", "Port to run yagg")
	host        = flag.String("b", "localhost", "Hostname to be used")
  repoPath    = flag.String("r", "data", "Set the git repositories path where data will be saved")
)

type Page struct {
	Title string
	Body  string
}

func init() {
  http.HandleFunc("/", NewHandler)
  http.HandleFunc("/create", CreateHandler)
}

func main() {
  flag.Parse()

  repo, err := CreateRepository(); if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Repository working directory set as:", repo.Workdir())

  addr := fmt.Sprintf("%s:%s", *host, *port)
  fmt.Println("Listening on ", addr)

  if err = http.ListenAndServe(addr, nil); err != nil {
    log.Fatal(err)
  }
}
