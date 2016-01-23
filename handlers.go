package main

import (
  "html/template"
  "net/http"
  "path"
  "io/ioutil"
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
  title := r.FormValue("gist-name")
  body := r.FormValue("gist")
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

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
  filename := r.URL.Path[len("/download/"):]
  file := path.Join(*repoPath, filename)
  data, err := ioutil.ReadFile(file); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Disposition", "attachment; filename="+filename)
  _, err = w.Write(data); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
