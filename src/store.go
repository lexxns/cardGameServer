package main

import (
	"encoding/json"
	"log"

	godux "github.com/luisvinicius167/godux"
)

var store *godux.Store

func initStore() {
	store = godux.NewStore()
}

func newRoom(roomID string) {
	initState := RoomState{
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
	store.SetState(roomID, initState)
}

// RoomState - roomState
type RoomState struct {
	Hand  []Card `json:"hand"`
	Field []Card `json:"field"`
}

// Card - card
type Card struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Attk   string `json:"attk"`
	Health string `json:"health"`
}

func roomStateMessage(roomID string) []byte {
	b, err := json.Marshal(store.GetState(roomID))

	if err != nil {
		log.Println(err)
		return nil
	}

	return b
}
