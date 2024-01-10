package main

import (
    "bytes"
    "html/template"
    "net/http"
)

var tasks []string
var tpl string = `
{{ define "index" }}
<!DOCTYPE html>
<html>
    <head>
        <script src="https://unpkg.com/htmx.org@1.9.10"></script>
        <title>Go/HTMX Tasks</title>
    </head>
    <body>
        <form hx-post="/tasks/add" hx-target="#task-list" hx-swap="innerHTML">
            <input name="title" type="text" placeholder="The title of the task">
            <button type="submit">Add task</button>
        </form>
        <ul id="task-list">
            {{ range .tasks }}
            <li>{{.}}</li>
            {{ else }}
            <li>No tasks found</li>
            {{ end }}
        </ul>
    </body>
</html>
{{ end }}
`

func addTask(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if len(r.Form.Get("title")) <= 0 {
        http.Error(w, "empty task title", http.StatusBadRequest)
        return
    }

    t, err := template.New("task-list").Parse(`
    {{ define "task-list" }}
        {{ range .tasks }}
        <li>{{.}}</li>
        {{ end }}
    {{ end }}
    `)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tasks = append(tasks, r.Form.Get("title"))

    var buf bytes.Buffer
    if err = t.Execute(&buf, map[string]interface{}{
        "tasks": tasks,
    }); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Write(buf.Bytes())
}

func index(w http.ResponseWriter, r *http.Request) {
    t, err := template.New("index").Parse(tpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var buf bytes.Buffer
    if err = t.Execute(&buf, map[string]interface{}{
        "tasks": tasks,
    }); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Write(buf.Bytes())   
}

func main() {
    http.HandleFunc("/", index) 
    http.HandleFunc("/tasks/add", addTask)
    http.ListenAndServe(":8080", nil)
}
