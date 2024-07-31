package database

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var schema = `
	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		creator text NOT NULL,
		recipient text NOT NULL,
		mail text NOT NULL,
		created date NOT NULL
	);

	ALTER TABLE messages REPLICA IDENTITY FULL;
`

type SqlConnection struct {
	Db *sqlx.DB
}

func NewConnection(connection string) (SqlConnection, error) {
	db, err := sqlx.Connect("pgx", connection)

	if err != nil {
		return SqlConnection{}, err
	}

	if err = db.Ping(); err != nil {
		return SqlConnection{}, err
	}

	db.MustExec(schema)

	return SqlConnection{Db: db}, nil
}
