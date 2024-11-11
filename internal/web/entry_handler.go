package web

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rafael-italiano/Pocket-Planner/internal/service"
)

type EntryHandlers struct {
	service *service.EntryService
}

func NewEntryHandlers(service *service.EntryService) *EntryHandlers {
	return &EntryHandlers{service: service}
}

func (h *EntryHandlers) GetEntries(w http.ResponseWriter, r *http.Request) {
	entries, err := h.service.GetEntries()
	if err != nil {
		http.Error(w, "failed to get entries", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}
func (h *EntryHandlers) CreateEntry(w http.ResponseWriter, r *http.Request) {
	var entry service.Entry
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.CreateEntry(&entry)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entry)
}
