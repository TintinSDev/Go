package main

import (
    "fmt"
    "net/http"
)

// HomePage handler
func HomePage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func main() {
    http.HandleFunc("/", HomePage)
    fmt.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
