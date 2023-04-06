package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/models"
)

func (c *Client) Login(w http.ResponseWriter, r *http.Request) {
	var credintails models.Credintails
	err := json.NewDecoder(r.Body).Decode(&credintails)
	if err != nil {
		c.logger.Error.Print(err)
		return
	}

	c.logger.Debug.Print(credintails)
}

func (c *Client) Register(w http.ResponseWriter, r *http.Request) {
}
