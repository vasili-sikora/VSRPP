package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"lab4/database"

	_ "github.com/mattn/go-sqlite3"
)

type AddRequest struct {
	Text string `json:"text"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func addHandler(db *database.SQLiteDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req AddRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}

		if req.Text == "" {
			http.Error(w, "text is empty", http.StatusBadRequest)
			return
		}

		if err := db.Insert(req.Text); err != nil {
			http.Error(w, "failed to insert into db", http.StatusInternalServerError)
			return
		}

		WriteJSON(w, http.StatusCreated, map[string]string{
			"message": "saved",
		})
	}
}

func allHandler(db *database.SQLiteDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		data, err := db.GetAll()
		if err != nil {
			http.Error(w, "failed to get data from db", http.StatusInternalServerError)
			return
		}

		WriteJSON(w, http.StatusOK, data)
	}
}

func main() {
	conn, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	db := database.New(conn)

	if err := db.CreateTable(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/add", addHandler(db))
	mux.HandleFunc("/all", allHandler(db))

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
