package repository

import (
	"database/sql"
	"os"

	"github.com/expose443/real-time-forum-golang/backend-api/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDB(cfg *config.DB) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.Dbname)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	err = createTables(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTables(db *sql.DB) error {
	path := "./internal/repository/tables/"
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, v := range dir {
		query, err := os.ReadFile(path + v.Name())
		if err != nil {
			return err
		}
		_, err = db.Exec(string(query))
		if err != nil {
			return err
		}

	}
	return nil
}
