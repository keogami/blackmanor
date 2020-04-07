package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"io/ioutil"
	"github.com/gorilla/websocket"
	"github.com/keogami/blackmanor/player"
)

var (
	totalConns int = 0
	player [2]*player.Player
)

func startExchange() {
	var m player.Message
	for {
		select {
		case m, ok := <- player[0].Out:
			if !ok {
				break
			}
			player[1].in <- m
		case m, ok := <- player[1].Out:
			if !ok {
				break
			}
			player[0].in <- m
		}
	}
}

func shiftContext(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	if totalConns == 0 {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.WriteMessage(websocket.TextMessage, "Waiting for the other Player to arrive")
		player[0] = player.New(0, conn)
		go player[0].ListenRead()
		go player[0].SendWrites()
		totalConns  += 1
	} else if totalConns == 1 {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		player.Message = Message{ websocket.TextMessage, "You may start sending messages now" }
		player[1] = player.New(1, conn)
		go player[1].ListenRead()
		go player[1].SendWrites()
		go startExchange()
		totalConns  += 1
	} else {
		fmt.FPrintf(w, "Connection refused.")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	mappedPath := ""
	if r.URL.Path == "/" {
		mappedPath = "index.html"
	} else if r.URL.Path[1:] == "socketgateway" {
		shiftContext(w, r)
		return
	}  else {
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