package storage

import "database/sql"

func CreateUserTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS USER (
		id SERIAL PRIMARY KEY,
		contact BIGINT,
		email VARCHAR(60),
		password VARCHAR(60),
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	)`

	_, err := db.Exec(query)

	return err
}
