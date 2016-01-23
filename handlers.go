package main

import (
  "html/template"
  "net/http"
  "path"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
  file := path.Join("view", "new.html")
  t, err := template.ParseFiles(file)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = t.Execute(w, nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
