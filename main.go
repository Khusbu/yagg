package main

import (
  "flag"
  "fmt"
  "net/http"
  "log"
)

var (
	port       = flag.String("p", "8080", "Port to run yagg")
	host       = flag.String("b", "localhost", "Hostname to be used")
)

func main() {
  flag.Parse()

  addr := fmt.Sprintf("%s:%s", *host, *port)
  if err := http.ListenAndServe(addr, nil); err != nil{
    log.Fatal(err)
  }
}
