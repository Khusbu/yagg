package main

import (
  "html/template"
  "net/http"
  "path"
  "io/ioutil"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    t, err := template.ParseFiles(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
  file := path.Join("view", "new.html")
  renderTemplate(w, file, nil)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
  title := r.FormValue("gist-name")
  body := r.FormValue("gist")
  api := r.FormValue("api")
  file_list, err := GetFileList(*repoPath); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  if api == "new" && !CheckFileName(file_list, title) {
    file := path.Join("view", "new.html")
    renderTemplate(w, file, &Error{Message: "Gist Name already exists"})
  } else {
    payload := &Page{Title: title, Body: []byte(body)}
    if err := payload.Save(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/show/" + title, http.StatusFound)
  }
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/show/"):]
  file := path.Join("view", "show.html")
  payload, err := GetPayload(title); if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }
  renderTemplate(w, file, payload)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/edit/"):]
  file := path.Join("view", "edit.html")
  payload, err := GetPayload(title); if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }
  renderTemplate(w, file, payload)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
  filename := r.URL.Path[len("/download/"):]
  file := path.Join(*repoPath, filename)
  data, err := ioutil.ReadFile(file); if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }
  w.Header().Set("Content-Disposition", "attachment; filename="+filename)
  _, err = w.Write(data); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func RawHandler(w http.ResponseWriter, r *http.Request) {
  filename := r.URL.Path[len("/raw/"):]
  file := path.Join(*repoPath, filename)
  data, err := ioutil.ReadFile(file); if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }
  _, err = w.Write(data); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  file := path.Join("view", "index.html")
  file_list, err := GetFileList(*repoPath); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  f := &Files{List: file_list};
  renderTemplate(w, file, f)
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/history/"):]
  history, err := GetHistory(title); if err != nil {
      http.Error(w, err.Error(), http.StatusNotFound)
      return
  }
  file := path.Join("view", "history.html")
  renderTemplate(w, file, history)
}

func ShowByIdHandler(w http.ResponseWriter, r *http.Request) {
  filename, rawId := GetFileAndRawId(r.URL.Path, "/show-by-id/")
  data, err := GetData(rawId); if err != nil {
      http.Error(w, err.Error(), http.StatusNotFound)
      return
  }
  file := path.Join("view", "show.html")
  renderTemplate(w, file, &Page{Title: filename, Body: data})
}

func RawByIdHandler(w http.ResponseWriter, r *http.Request) {
  _, rawId := GetFileAndRawId(r.URL.Path, "/raw-by-id/")
  data, err := GetData(rawId); if err != nil {
      http.Error(w, err.Error(), http.StatusNotFound)
      return
  }
  _, err = w.Write(data); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func DownloadByIdHandler(w http.ResponseWriter, r *http.Request) {
  filename, rawId := GetFileAndRawId(r.URL.Path, "/download-by-id/")
  data, err := GetData(rawId); if err != nil {
      http.Error(w, err.Error(), http.StatusNotFound)
      return
  }
  w.Header().Set("Content-Disposition", "attachment; filename="+filename)
  _, err = w.Write(data); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
