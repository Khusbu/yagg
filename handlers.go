package main

import (
  "html/template"
  "net/http"
  "path"
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
  http.Redirect(w, r, "/show/" + title, http.StatusFound)
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/show/"):]
  file := path.Join("view", "show.html")
  t, err := template.ParseFiles(file); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  payload, err := GetPayload(title); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = t.Execute(w, payload); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/edit/"):]
  file := path.Join("view", "edit.html")
  t, err := template.ParseFiles(file); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  payload, err := GetPayload(title); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = t.Execute(w, payload); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
