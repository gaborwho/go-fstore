package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var fsHandler_Implements http.Handler = (*fsHandler)(nil)

type fsHandler struct {
}

func (vh *fsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := "./" + r.RequestURI
	log.Println(path)
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		log.Println("NotExist")
		return
	}
	check(err)

	fi, err := file.Stat()
	check(err)

	if !fi.IsDir() {
		dat, err := ioutil.ReadFile("." + r.RequestURI)
		check(err)
		w.Write(dat)
	}

}
