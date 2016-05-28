package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"tanzania-website/messages"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/messages", messages.MainController)
	http.HandleFunc("/", serveResource)

	port := "8080"

	fmt.Println("Server started on port " + port)
	http.ListenAndServe(":"+port, nil)
}

func serveResource(w http.ResponseWriter, req *http.Request) {

	path := "public" + req.URL.Path

	var contentType string

	if path == "public/" {
		path = "public/introduction.html"
	}

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg"
	} else if strings.HasSuffix(path, ".html") {
		contentType = "text/html"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
