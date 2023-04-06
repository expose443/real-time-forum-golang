package repository

import (
	"github.com/expose443/real-time-forum-golang/backend-api/internal/models"
)

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users(nickname, age, gender, first_name, last_name, email, password, created, updated) values(?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password, user.Created, user.Updated)
	return err
}

func (r *UserRepository) GetUserByIdentifier(identifier string) (models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = ? OR nickname = ?`
	err := r.db.QueryRow(query, identifier, identifier).Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Created, &user.Updated)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByUserID(id int) (models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE user_id = ?`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Created, &user.Updated)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
