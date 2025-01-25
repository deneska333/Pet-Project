package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitDB()
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", PostHandler).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	DB.Find(&messages)
	json.NewEncoder(w).Encode(messages)
}

type MessageRequest struct {
	Message string `json:"message"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var message MessageRequest

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	InsertTask(message)

	log.Println(message)
}
