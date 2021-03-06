package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("listening...")
    err := http.ListenAndServe(getPort(), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello. This is our first Go web app on Heroku!")
}

func getPort() string {
    port := os.Getenv("PORT")
    if port == "" {
        port = "4747"
        fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
    }
    return ":" + port
}
