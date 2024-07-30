package database

import (
	"github.com/AlexEr256/messageService/environments"
	"github.com/jmoiron/sqlx"
)

type SqlConnection struct {
	Db *sqlx.DB
}

func NewConnection() (SqlConnection, error) {
	db, err := sqlx.Connect("postgres", environments.ConnectionString)
	if err != nil {
		return SqlConnection{}, err
	}

	if err = db.Ping(); err != nil {
		return SqlConnection{}, err
	}

	return SqlConnection{Db: db}, nil
}
