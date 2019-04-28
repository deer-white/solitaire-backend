package route

import (
	"encoding/json"
	"fmt"
	"net/http"
)

import (
	Solitaire "../solitaire"
)

func (r Router) HintHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	//strCards, ok := req.URL.Query()["cards"]

	strCards := req.FormValue("cards")

	if len(strCards) < 1 {
		fmt.Println("Url Param 'cards' is missing")
		w.WriteHeader(400)
		w.Write([]byte("need param"))
		return
	}

	var cards Solitaire.Cards
	err := json.Unmarshal([]byte(strCards), &cards)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(strCards)
	fmt.Println(cards)

	w.Write([]byte("ok"))
}