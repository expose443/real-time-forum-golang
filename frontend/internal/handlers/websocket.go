package handlers

import (
	"net/http"
	"text/template"
)

func WebSocket(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/html/index.html")
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
