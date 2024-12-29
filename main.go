package main

import (
    "fmt"
    "net/http"
    "os"
)

func fileHandler(w http.ResponseWriter, r *http.Request) {
    // Open the HTML file
    file, err := os.Open("index.html")
    if err != nil {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }
    defer file.Close()

    // Get file info
    fileInfo, err := file.Stat()
    if err != nil {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }

    // Set the Content-Type header
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    // Serve the file content
    http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}

func main() {
    http.HandleFunc("/", fileHandler)
    
    port := ":8080"
    fmt.Printf("Server is starting on port %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}
