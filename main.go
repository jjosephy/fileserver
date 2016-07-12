package main

import (
    "log"
    "net/http"
)

func main() {
    ///Users/q4vy/Source/go/src/github.com/jjosephy/server
    log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/Users/q4vy/Source/go/src/github.com/jjosephy/server"))))
}
