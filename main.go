package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
        name := r.PathValue("name")
        fmt.Fprintf(w, "Hello %s!", name)
    })
}
