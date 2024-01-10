package main

import (
    "fmt"
    "net/http"
)

var counter int
var tpl string = `
<!DOCTYPE html>
<html>
    <head>
        <script src="https://unpkg.com/htmx.org@1.9.10"></script>
        <title>Go/HTMX Counter</title>
    </head>
    </body>
        <p id="counter">0</p>
        <button hx-get="/increment" hx-target="#counter" hx-swap="outerHTML">Increment</button>
    </body>
</html>
`

func increment(w http.ResponseWriter, r *http.Request) {
    counter++
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Write([]byte(fmt.Sprintf(`<p id="counter">%d</p>`, counter)))
}

func index(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Write([]byte(tpl))   
}

func main() {
    http.HandleFunc("/", index) 
    http.HandleFunc("/increment", increment)
    http.ListenAndServe(":8080", nil)
}
