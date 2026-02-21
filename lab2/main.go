package main

import (
	"database/sql"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTable := `CREATE TABLE IF NOT EXISTS database(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			text TEXT NOT NULL
			);`
	_, err = db.Exec(createTable)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO database (text) VALUES (?)`, "hello fyne!")
	var message string
	err = db.QueryRow("SELECT text FROM database").Scan(&message)
	if err != nil {
		panic(err)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("sqlite3 test")

	label := widget.NewLabel(message)

	myWindow.SetContent(container.NewVBox(label))

	myWindow.ShowAndRun()

}
