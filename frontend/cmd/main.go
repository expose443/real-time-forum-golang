package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/ws", webSocket)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func webSocket(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("././templates/html/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = temp.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
