package main

import (
	"encoding/json"
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
		neffos.OnNamespaceConnected: onConnect,
		neffos.OnRoomJoined:         onRoomJoined,
	},
}

// RoomState - roomState
type RoomState struct {
	Hand  []Card `json:"hand"`
	Field []Card `json:"field"`
}

// Card - card
type Card struct {
	Name string `json:"name"`
}

//TODO Return current room state
func roomState() []byte {
	state := RoomState{
		Hand: []Card{
			{Name: "card 1"},
			{Name: "card 2"},
			{Name: "card 3"},
		},
		Field: []Card{
			{Name: "card 4"},
			{Name: "card 5"},
			{Name: "card 6"},
		},
	}
	b, err := json.Marshal(state)

	if err != nil {
		log.Println(err)
		return nil
	}

	return b
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
	c.Emit("state", roomState())
	return nil
}

func onConnect(c *neffos.NSConn, msg neffos.Message) error {
	//// TODO: Create a room or find an existing one
	roomName := []byte("some_room_name")
	c.Emit("room", roomName)
	return nil
}

func startServer() {
	websocketServer := neffos.New(upgrader, events)

	router := http.NewServeMux()
	router.Handle("/game", websocketServer)

	log.Println("Serving websockets on localhost:8080/game")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	startServer()
}
