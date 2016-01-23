package main

import (
  "html/template"
  "net/http"
  "path"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
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

func CreateHandler(w http.ResponseWriter, r *http.Request) {
    title := r.FormValue("paste-name")
    body := r.FormValue("paste")
    payload := &Page{Title: title, Body: body}
    file := path.Join("view", "create.html")
    t, err := template.ParseFiles(file)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    err = t.Execute(w, payload)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}
