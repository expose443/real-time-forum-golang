package repository

import "database/sql"

type Repository struct {
	*UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewUserRepository(db),
	}
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
