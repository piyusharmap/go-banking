package storage

import "database/sql"

func CreateUserTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		contact VARCHAR(15) UNIQUE NOT NULL,
		email VARCHAR(120) UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(query)

	return err
}

func CreateAccountTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
		first_name VARCHAR(50) NOT NULL,
		last_name VARCHAR(50),
		account_number VARCHAR(20) UNIQUE NOT NULL,
		balance BIGINT DEFAULT 000 NOT NULL,
		currency VARCHAR(3) NOT NULL DEFAULT 'INR',
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(query)

	return err
}
