package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey yo, nice to see you %s.", r.URL.Path[1:])
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