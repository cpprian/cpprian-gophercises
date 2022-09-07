package dbutil

const (
	Host    = "localhost"
	Port    = 5432
	User    = "postgres"
	Dbname  = "phone_normalizer"

	Create_database = `CREATE DATABASE phone_normalizer;`
	Create_table = `CREATE TABLE IF NOT EXISTS phone_numbers (
		id SERIAL PRIMARY KEY,
		phone_number VARCHAR(25) NOT NULL);`
	Delete_database = `DROP DATABASE IF EXISTS phone_normalizer;`
	Delete_table = `DROP TABLE IF EXISTS phone_numbers;`

	Insert_phone_number = `INSERT INTO phone_numbers (phone_number) VALUES ($1);`
	Select_phone_number = `SELECT id, phone_number FROM phone_numbers;`
	Delete_phone_number = `DELETE FROM phone_numbers WHERE phone_number = $1;`
)

