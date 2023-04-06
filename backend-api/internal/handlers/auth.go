package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/models"
)

func (c *Client) Login(w http.ResponseWriter, r *http.Request) {
	var credintails models.Credintails
	err := json.NewDecoder(r.Body).Decode(&credintails)
	if err != nil {
		c.Logger.Error.Print(err)
		return
	}
	user, err := c.AuthService.IsValidUser(credintails.Identifier, credintails.Password)
	if err != nil {
		c.Logger.Error.Print(err)
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "hello %s", user.FirstName)
}

func (c *Client) Register(w http.ResponseWriter, r *http.Request) {
	var user models.UserRegister
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		c.Logger.Error.Print(err)
		return
	}
	if user.Password != user.ConfirmPassword {
		c.Logger.Debug.Print("password doesn't match")
		w.WriteHeader(400)
		return
	}

	err = c.AuthService.CreateUser(&models.User{
		Nickname:  user.Nickname,
		Age:       user.Age,
		Gender:    user.Gender,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	})
	if err != nil {
		c.Logger.Error.Print(err)
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
}
