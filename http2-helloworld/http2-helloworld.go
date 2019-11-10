package main

import (
	"io"
	"net/http"
	"log"
	"flag"
	"fmt"

	"golang.org/x/net/http2"
)

func main() {
	const indexHTML = `<html><head><title>Hello World</title><script src="/static/app.js"></script><link rel="stylesheet" href="/static/style.css"></head><body>Hello, gopher!</body></html>`

	var srv http.Server
	srv.Addr = ":8080"
	http2.ConfigureServer(&srv, nil)
	flag.Parse()
	
	http.Handle("/static/",
			http.StripPrefix("/static/",
					http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		pusher, ok := w.(http.Pusher)
		if ok {
			err := pusher.Push("/static/app.js", nil)
			if err != nil {
				log.Printf("Failed to push: %v", err)
			}
			err = pusher.Push("/static/style.css", nil)
			if err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		fmt.Fprintf(w, indexHTML)
	})

	err := srv.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
        log.Printf("[ERROR] %s", err)
    }
}
func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World! by HTTP/2")
}