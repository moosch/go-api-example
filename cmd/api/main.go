package main

import (
    "fmt"
    "net/http"

    log "github.com/sirupsen/logrus"

    "github.com/moosch/go-api-example/internal/handlers"
    "github.com/moosch/go-api-example/internal/cache"
)

func main() {
    log.SetReportCaller(true)
    mux := http.NewServeMux()

    cache.Init()

    handlers.Handler(mux)

    fmt.Println("Starting Go API...")

    err := http.ListenAndServe(":8000", mux)
    if err != nil {
        log.Error(err)
    }
}

