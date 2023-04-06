package repository

import "database/sql"

type DAO interface {
	NewUserRepo() UserRepository
}

type dao struct {
	db *sql.DB
}

func NewDao(db *sql.DB) DAO {
	return &dao{
		db: db,
	}
}

func (d dao) NewUserRepo() UserRepository {
	return &userRepository{
		db: d.db,
	}
}
