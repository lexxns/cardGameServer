package main

import (
	"encoding/json"
	"log"

	sid "github.com/chilts/sid"
	godux "github.com/luisvinicius167/godux"
)

var store *godux.Store

func initStore() {
	store = godux.NewStore()
}

func newRoom(roomID string) {
	card1 := Card{ID: sid.Id(), Name: "card 1", Attk: 1, Health: 1, Container: "HAND"}
	card2 := Card{ID: sid.Id(), Name: "card 2", Attk: 2, Health: 2, Container: "HAND"}
	card3 := Card{ID: sid.Id(), Name: "card 3", Attk: 3, Health: 3, Container: "HAND"}
	card4 := Card{ID: sid.Id(), Name: "card 4", Attk: 4, Health: 4, Container: "FIELD"}
	card5 := Card{ID: sid.Id(), Name: "card 5", Attk: 5, Health: 5, Container: "FIELD"}
	card6 := Card{ID: sid.Id(), Name: "card 6", Attk: 6, Health: 6, Container: "FIELD"}
	initState := RoomState{
		Cards: map[string]Card{
			card1.ID: card1,
			card2.ID: card2,
			card3.ID: card3,
			card4.ID: card4,
			card5.ID: card5,
			card6.ID: card6,
		},
	}
	store.SetState(roomID, initState)
}

// RoomState - roomState
type RoomState struct {
	Cards map[string]Card `json:"cards"`
}

// Card - card
type Card struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Attk      int    `json:"attk"`
	Health    int    `json:"health"`
	Container string `json:"container"`
}

func incAttk(number int) godux.Action {
	return godux.Action{
		Type:  "INC_ATTK",
		Value: number,
	}
}

func reducer(roomID string, cardID string, action godux.Action) interface{} {
	switch action.Type {
	case "INC_ATTK":
		roomState := store.GetState(roomID).(RoomState)
		var card = roomState.Cards[cardID]
		card.Attk += action.Value.(int)
		roomState.Cards[cardID] = card
		return roomState
	default:
		return store.GetAllState()
	}

}

func roomStateMessage(roomID string) []byte {
	b, err := json.Marshal(store.GetState(roomID))

	if err != nil {
		log.Println(err)
		return nil
	}

	return b
}
