package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "serve dir")
	var addr string
	flag.StringVar(&addr, "addr", ":8081", "serve addr")
	flag.Parse()

	log.Printf("addr=%s", addr)
	log.Printf("dir=%s", dir)

	http.Handle("/", http.FileServer(http.Dir(dir)))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
