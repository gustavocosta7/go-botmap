package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	log.Println("R: ")
	log.Println(r)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("CLient connected")

	err = ws.WriteMessage(1, []byte("Hi client"))

	if err != nil {
		log.Println(" - Erro -")
		log.Println(err)
	}
	
	reader(ws)
}

func reader(conn *websocket.Conn) {
	for  {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
		}

		if string(p) == "oi" {
			if err := conn.WriteMessage(messageType, []byte("Olá pessoas")); err != nil {
				log.Println(err)
				return
			}
		}


		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}


func setupRoutes()  {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello world")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
