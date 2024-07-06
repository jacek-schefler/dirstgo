package main

import (
	"log"

	"github.com/jacek-schefler/distrgo/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
