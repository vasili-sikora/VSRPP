package database

import "database/sql"

type SQLiteDB struct {
	conn *sql.DB
}

func New(conn *sql.DB) *SQLiteDB {
	return &SQLiteDB{conn: conn}
}

func (d *SQLiteDB) CreateTable() error {
	_, err := d.conn.Exec(`CREATE TABLE IF NOT EXISTS database(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text NOT NULL
		)`)

	return err
}

func (d *SQLiteDB) Insert(text string) error {
	_, err := d.conn.Exec(`INSERT INTO database (text) VALUES (?)`, text)
	return err
}

func (d *SQLiteDB) GetFirst() (string, error) {
	var message string
	err := d.conn.QueryRow("SELECT text from database").Scan(&message)
	return message, err
}
