package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rafael-italiano/Pocket-Planner/internal/service"
	"github.com/rafael-italiano/Pocket-Planner/internal/web"
)

func main() {
	db, err := sql.Open("sqlite3", "./database/ledger.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	entryService := service.NewEntryService(db)
	entryHandler := web.NewEntryHandlers(entryService)

	router := http.NewServeMux()

	router.HandleFunc("GET /entries", entryHandler.GetEntries)
	router.HandleFunc("POST /entries", entryHandler.CreateEntry)
	/*router.HandleFunc("GET /books/{id}", entryHandler.GetBookByID)
	router.HandleFunc("PUT /books/{id}", entryHandler.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", entryHandler.DeleteBook)*/
	log.Println("Starting server...")
	http.ListenAndServe(":8080", router)
}
