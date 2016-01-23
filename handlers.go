package main

import (
  "html/template"
  "net/http"
  "path"
  "fmt"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
  file := path.Join("view", "new.html")
  t, err := template.ParseFiles(file); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = t.Execute(w, nil); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
  title := r.FormValue("paste-name")
  body := r.FormValue("paste")
  payload := &Page{Title: title, Body: []byte(body)}
  if err := payload.save(); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
  http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  file := path.Join(*repoPath, title)
  http.ServeFile(w, r, file)
}
