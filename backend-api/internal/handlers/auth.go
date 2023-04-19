package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/jwt"
	"github.com/expose443/real-time-forum-golang/backend-api/internal/models"
)

func (c *Client) Login(w http.ResponseWriter, r *http.Request) {
	var credintails models.Credintails
	err := json.NewDecoder(r.Body).Decode(&credintails)
	if err != nil {
		c.logger.Error.Print(err)
		return
	}
	user, err := c.authService.IsValidUser(credintails.Identifier, credintails.Password)
	if err != nil {
		c.logger.Error.Print(err)
		w.WriteHeader(400)
		return
	}

	expiry := time.Now().Add(time.Second * 10)
	expiryStr := expiry.Format(time.RFC3339)
	claims := map[string]interface{}{
		"exp": expiryStr,
		"sub": user.ID,
	}

	jwtToken, err := jwt.CreateJWT(claims)
	if err != nil {
		c.logger.Error.Print(err)
		w.WriteHeader(500)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Value:   jwtToken,
		Name:    "jwt_token",
		Expires: expiry,
	})

	w.WriteHeader(200)
	fmt.Fprintf(w, "hello token: %s", jwtToken)
}

func (c *Client) Register(w http.ResponseWriter, r *http.Request) {
	var user models.UserRegister
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		c.logger.Error.Print(err)
		return
	}
	if user.Password != user.ConfirmPassword {
		c.logger.Debug.Print("password doesn't match")
		w.WriteHeader(400)
		return
	}

	err = c.authService.CreateUser(&models.User{
		Nickname:  user.Nickname,
		Age:       user.Age,
		Gender:    user.Gender,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	})
	if err != nil {
		c.logger.Error.Print(err)
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
}
