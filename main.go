package main

import (
	"fmt"
	"net/http"
)

import (
	Route "./route"
)

func main() {
	port := ":5000"
	fmt.Println("Start http serve at " + port)

	var router Route.Router

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", router.NewGameHandler)
	http.HandleFunc("/hint", router.HintHandler)

	http.ListenAndServe(port, nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hi there"))
}


