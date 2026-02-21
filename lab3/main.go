package main

import (
	"database/sql"
	"log"

	"lab3/app"
	"lab3/database"

	fyneapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	database := database.New(conn)
	app := app.New(database)

	message, err := app.Run()
	if err != nil {
		log.Fatal(err)
	}

	myApp := fyneapp.New()
	myWindow := myApp.NewWindow("sqlite3 test")
	myWindow.SetContent(container.NewVBox(widget.NewLabel(message)))
	myWindow.ShowAndRun()
}
