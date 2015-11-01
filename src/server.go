package main

import (
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	handler := &fsHandler{}
	log.Fatal(http.ListenAndServe(":8082", handler))
}
