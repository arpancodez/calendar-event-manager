package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// Event represents a calendar event
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Location    string    `json:"location"`
}

// EventStore manages events in memory
type EventStore struct {
	mu     sync.RWMutex
	events map[string]Event
}

var store = &EventStore{
	events: make(map[string]Event),
}

// CreateEvent handles POST requests to create a new event
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate simple ID
	event.ID = fmt.Sprintf("%d", time.Now().Unix())

	store.mu.Lock()
	store.events[event.ID] = event
	store.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

// GetAllEvents handles GET requests to retrieve all events
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	events := make([]Event, 0, len(store.events))
	for _, event := range store.events {
		events = append(events, event)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// GetEvent handles GET requests to retrieve a specific event
func GetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	store.mu.RLock()
	event, exists := store.events[id]
	store.mu.RUnlock()

	if !exists {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

// UpdateEvent handles PUT requests to update an event
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.mu.Lock()
	if _, exists := store.events[id]; !exists {
		store.mu.Unlock()
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}
	event.ID = id
	store.events[id] = event
	store.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

// DeleteEvent handles DELETE requests to remove an event
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	store.mu.Lock()
	if _, exists := store.events[id]; !exists {
		store.mu.Unlock()
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}
	delete(store.events, id)
	store.mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

// ServeHTML serves the frontend HTML file
func ServeHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	// API routes
	router.HandleFunc("/api/events", CreateEvent).Methods("POST")
	router.HandleFunc("/api/events", GetAllEvents).Methods("GET")
	router.HandleFunc("/api/events/{id}", GetEvent).Methods("GET")
	router.HandleFunc("/api/events/{id}", UpdateEvent).Methods("PUT")
	router.HandleFunc("/api/events/{id}", DeleteEvent).Methods("DELETE")

	// Serve frontend
	router.HandleFunc("/", ServeHTML).Methods("GET")

	// Apply CORS middleware
	handler := enableCORS(router)

	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
