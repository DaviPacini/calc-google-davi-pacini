package main

import (
	"log"
	"net/http"
)

func StartServer() {

	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: JapaHandler{},
		//ReadTimeout:	10 * time.Second,
		//WhriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
