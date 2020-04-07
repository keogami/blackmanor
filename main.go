package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	mappedPath := ""
	if r.URL.Path == "/" {
		mappedPath = "index.html"
	} else {
		mappedPath = "." + r.URL.Path
	}
	content, err := ioutil.ReadFile(mappedPath)
	if err != nil {
		http.Error(w, "this file doesn't not exist", http.StatusNotFound)
	}
	fmt.Fprintf(w, "%s", string(content))
	fmt.Println(r.URL.Path)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT is not set in the enviroment")
	}
	http.HandleFunc("/", handler)
	fmt.Println("Attempting to run on $PORT=" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}