package e2e

import (
	"database/sql"
	"log"
)

func createTestTables(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			balance INT NOT NULL DEFAULT 1000
		);

		CREATE TABLE IF NOT EXISTS merch (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE NOT NULL,
			price INT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS purchases (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id),
			merch_id INT REFERENCES merch(id),
			price_bought INT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		);

		CREATE TABLE IF NOT EXISTS coin_transfers (
			id SERIAL PRIMARY KEY,
			from_user_id INT REFERENCES users(id),
			to_user_id INT REFERENCES users(id),
			amount INT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		);

		INSERT INTO merch (name, price) VALUES ('test-item', 100) ON CONFLICT (name) DO NOTHING;
	`)
	if err != nil {
		log.Fatalf("Failed to create test tables: %s", err)
	}
}

func cleanupTestTables(db *sql.DB) {
	_, err := db.Exec(`
		DROP TABLE IF EXISTS coin_transfers CASCADE;
		DROP TABLE IF EXISTS purchases CASCADE;
		DROP TABLE IF EXISTS merch CASCADE;
		DROP TABLE IF EXISTS users CASCADE;
	`)
	if err != nil {
		log.Fatalf("Failed to clean up test tables: %s", err)
	}
}
