package repository

import (
	"database/sql"
	"time"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByIdentifier(identifier string) (models.User, error)
	GetUserByUserID(id int) (models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users(nickname, age, gender, first_name, last_name, email, password, created, updated) values(?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := u.db.Exec(query, user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password, time.Now(), time.Now())
	return err
}

func (u *userRepository) GetUserByIdentifier(identifier string) (models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = ? OR nickname = ?`
	err := u.db.QueryRow(query, identifier, identifier).Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Created, &user.Updated)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *userRepository) GetUserByUserID(id int) (models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE user_id = ?`
	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Created, &user.Updated)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
