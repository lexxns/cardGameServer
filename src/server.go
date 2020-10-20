package main

import (
	"log"
	"net/http"

	sid "github.com/chilts/sid"
	"github.com/gorilla/websocket"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
)

var upgrader = gorilla.Upgrader(websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
})

var events = neffos.Namespaces{
	"v1": neffos.Events{
		neffos.OnNamespaceConnected: onConnect,
		neffos.OnRoomJoined:         onRoomJoined,
	},
}

func createMessage(msg neffos.Message, event string, body []byte) neffos.Message {
	return neffos.Message{
		Namespace: msg.Namespace,
		Room:      msg.Room,
		Event:     event,
		Body:      body,
	}
}

func onRoomJoined(c *neffos.NSConn, msg neffos.Message) error {
	newRoom(msg.Room)
	c.Emit("state", roomStateMessage(msg.Room))
	return nil
}

func onConnect(c *neffos.NSConn, msg neffos.Message) error {
	//// TODO: Find a way of getting room id's to clients
	//// So existing rooms can be joined
	roomName := []byte(sid.Id())
	c.Emit("room", roomName)
	return nil
}

func startServer() {
	websocketServer := neffos.New(upgrader, events)

	router := http.NewServeMux()
	router.Handle("/game", websocketServer)

	initStore()

	log.Println("Serving websockets on localhost:8080/game")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	startServer()
}
