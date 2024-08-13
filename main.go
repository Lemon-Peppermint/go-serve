package main

import (
	//"database/sql"
	"log"
	"net/http"
	//"os"
  "io/ioutil"

  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"
)

func main() {
	logger := log.Default()
	logger.SetFlags(3)

  router := chi.NewRouter()
  router.Use(middleware.RequestID)
  router.Use(middleware.RealIP)
  router.Use(middleware.Logger)
  router.Use(middleware.Recoverer)

  router.Get("/", func (w http.ResponseWriter, r *http.Request) {
    file, err := ioutil.ReadFile("./static/index.html")

    if err != nil {
      log.Fatal("Couldn't read file: ./static/index.html")
    }

    w.Write(file)
  })

  router.Method("Get", "/static", http.FileServer(http.Dir("./static/")))

	log.Println("Hello serve")

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

