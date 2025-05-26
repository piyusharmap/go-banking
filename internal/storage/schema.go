package storage

import "database/sql"

func CreateUserTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		usr_id SERIAL PRIMARY KEY,
		contact BIGINT,
		email VARCHAR(60),
		password VARCHAR(60),
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	)`

	_, err := db.Exec(query)

	return err
}

func CreateAccountTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS account (
		acc_id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		account_number VARCHAR(20),
		balance BIGINT,
		created_at TIMESTAMP,
		updated_at TIMESTAMP,
		CONSTRAINT fk_users FOREIGN KEY (usr_id)
		REFERENCES users(usr_id)
	)`

	_, err := db.Exec(query)

	return err
}
