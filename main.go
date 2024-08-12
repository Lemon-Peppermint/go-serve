package main

import (
	//"database/sql"
	//"fmt"
	"log"
	//"net"
	"net/http"
	//"os"
)

func main() {
	logger := log.Default()
	logger.SetFlags(3)

	log.Println("Hello serve")

	log.Fatal(http.ListenAndServe("localhost:8000", http.FileServer(http.Dir("./static"))))
}
