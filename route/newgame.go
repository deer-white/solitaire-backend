package route

import (
	"net/http"
)

func (r Router) NewGameHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	boardTemplate, ok := req.URL.Query()["template"]

	if !ok || len(boardTemplate[0]) < 1 {
		boardTemplate[0] = "default"
	}



	w.Write([]byte("ok"))
}