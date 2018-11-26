package main

import (
  "log"
  "net/http"
  "os"
  "path/filepath"
  "runtime"

  "shrt1/handlers"
  "shrt1/storages"
  "github.com/mitchellh/go-homedir"
)

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  dir, _ := homedir.Dir()
  storage := &storages.Filesystem{}
  err := storage.Init(filepath.Join(dir, "shrt1"))
  if err != nil {
    log.Fatal(err)
  }

  http.Handle("/", handlers.EncodeHandler(storage))
  http.Handle("/dec/", handlers.DecodeHandler(storage))
  http.Handle("/red/", handlers.RedirectHandler(storage))

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  log.Println("Listing on :" + port)

  err = http.ListenAndServe(":"+port, nil)
  if err != nil {
    log.Fatal(err)
  }
}

