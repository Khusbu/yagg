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
	Body  []byte
}

type Files struct {
  List []string
}

type History struct {
  Title string
  CDiffs []CommitDiff
}

func init() {
  err := CreateRepository(); if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Repository working directory set as:", repo.Workdir())
  http.HandleFunc("/", NewHandler)
  http.HandleFunc("/create", CreateHandler)
  http.HandleFunc("/show/", ShowHandler)
  http.HandleFunc("/edit/", EditHandler)
  http.HandleFunc("/download/", DownloadHandler)
  http.HandleFunc("/raw/", RawHandler)
  http.HandleFunc("/list/", IndexHandler)
  http.HandleFunc("/history/", HistoryHandler)
  http.HandleFunc("/show-by-id/", ShowByIdHandler)
  http.HandleFunc("/raw-by-id/", RawByIdHandler)
  http.HandleFunc("/download-by-id/", DownloadByIdHandler)
  http.Handle("/assets/",http.FileServer(http.Dir(".")) )
  http.Handle("/javascripts/",http.FileServer(http.Dir(".")) )
}

func main() {
  flag.Parse()

  addr := fmt.Sprintf("%s:%s", *host, *port)
  fmt.Println("Listening on ", addr)

  if err := http.ListenAndServe(addr, nil); err != nil {
    log.Fatal(err)
  }
}
