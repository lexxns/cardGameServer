package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
)

var upgrader = gorilla.Upgrader(websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
})

var events = neffos.Namespaces{
	"v1": neffos.Events{
		"echo": onEcho,
	},
}

func onEcho(c *neffos.NSConn, msg neffos.Message) error {
	body := string(msg.Body)
	log.Println(body)

	newBody := append([]byte("echo back: "), msg.Body...)
	return neffos.Reply(newBody)
}

func startServer() {
	websocketServer := neffos.New(upgrader, events)

	router := http.NewServeMux()
	router.Handle("/echo", websocketServer)

	log.Println("Serving websockets on localhost:8080/echo")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	startServer()
}
