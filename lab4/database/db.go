package database

import (
	"database/sql"
)

type SQLiteDB struct {
	conn *sql.DB
}

func New(conn *sql.DB) *SQLiteDB {
	return &SQLiteDB{conn: conn}
}

func (d *SQLiteDB) CreateTable() error {
	_, err := d.conn.Exec(`CREATE TABLE IF NOT EXISTS database(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL
		)`)

	return err
}

func (d *SQLiteDB) Insert(text string) error {
	_, err := d.conn.Exec(`INSERT INTO database (text) VALUES (?)`, text)
	return err
}

func (d *SQLiteDB) GetFirst() (string, error) {
	var message string
	err := d.conn.QueryRow("SELECT text FROM database").Scan(&message)
	return message, err
}

func (d *SQLiteDB) GetAll() ([]string, error) {
	rows, err := d.conn.Query("SELECT text FROM database")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []string
	for rows.Next() {
		var t string
		if err := rows.Scan(&t); err != nil {
			return nil, err
		}
		data = append(data, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
