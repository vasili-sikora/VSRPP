package main

import (
	"encoding/json"
	"errors"
	"lab4/database"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AddRequest struct {
	text string "json:text"
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func main() {
	conn, err := sql.Open("sqlite3", "database.db")
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()

	db:= database.New(conn)

	if err != db.CreateTable(); err != nil{
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost{
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

		var req addRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		if req.Text == ""{
			http.Error(w, "text is empty", http.StatusBadRequest)
			return
		}
		if err!=db.Insert(req.Text); err != nil{
			log.fatal(err)
		}
	}
}
