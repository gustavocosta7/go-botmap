package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go-websocket-connection/functions"
	"log"
	"net/http"
	"strconv"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {


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

		indexMessage,err := strconv.Atoi(string(p))
		if err != nil {
			fmt.Println("Erro ao converter numero")
		}

		var message = functions.GetMessage(indexMessage)
		if err := conn.WriteMessage(messageType, []byte(message)); err != nil {
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
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
