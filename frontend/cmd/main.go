package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Credintails struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type UserRegister struct {
	Nickname        string `json:"nickname"`
	Age             int    `json:"age"`
	Gender          string `json:"gender"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func main() {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./templates/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fs))

	router.HandleFunc("/ws", webSocket)
	router.HandleFunc("/sign-in", Signin)
	router.HandleFunc("/sign-up", Signup)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func webSocket(w http.ResponseWriter, r *http.Request) {
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

func Signin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("./templates/html/sign-in.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	case http.MethodPost:
		r.ParseForm()
		credintails := Credintails{
			Identifier: r.FormValue("identity"),
			Password:   r.FormValue("password"),
		}
		body, err := json.Marshal(&credintails)
		if err != nil {
			fmt.Println(err.Error())
		}
		req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/login", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println(err)
		}
		defer req.Body.Close()
		client := http.Client{
			Timeout: 10 * time.Second,
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		if res.StatusCode != 200 {
			w.WriteHeader(res.StatusCode)
			return
		}
		token := res.Cookies()

		http.SetCookie(w, &http.Cookie{
			Expires: token[0].Expires,
			Value:   token[0].Value,
			Name:    "jwt_token",
		})

	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("./templates/html/sign-up.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	case http.MethodPost:
		r.ParseForm()

		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			fmt.Println(err)
		}
		credintails := UserRegister{
			Nickname:        r.FormValue("nickname"),
			Age:             age,
			Gender:          r.FormValue("gender"),
			FirstName:       r.FormValue("first_name"),
			LastName:        r.FormValue("last_name"),
			Email:           r.FormValue("email"),
			Password:        r.FormValue("password"),
			ConfirmPassword: r.FormValue("confirm"),
		}
		body, err := json.Marshal(&credintails)
		if err != nil {
			fmt.Println(err.Error())
		}
		req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/register", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println(err)
		}

		client := http.Client{
			Timeout: 10 * time.Second,
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		if res.StatusCode != 200 {
			w.WriteHeader(res.StatusCode)
			return
		}
		// identity := r.FormValue("identity")
		// password := r.FormValue("password")

	}
}
