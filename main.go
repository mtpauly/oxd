package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    
    port := ":8080"
    fmt.Printf("Server is starting on port %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}