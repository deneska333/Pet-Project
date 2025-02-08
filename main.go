package main

import (
	repository "Project/Repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	repository.InitDB()
	repository.DB.AutoMigrate(&repository.Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", PostHandler).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/messages/{id}", PatchMessages).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", DeleteMessages).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []repository.Message
	repository.DB.Find(&messages)
	json.NewEncoder(w).Encode(messages)
	log.Println(messages)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var message repository.MessageRequest

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repository.InsertTask(message)

	log.Println(message)
}

func PatchMessages(w http.ResponseWriter, r *http.Request) {
	sId := mux.Vars(r)
	sid := sId["id"]
	id, _ := strconv.Atoi(sid)
	fmt.Println(id)

	var message repository.MessageRequest

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	repository.UpdateTask(id, message)
}

func DeleteMessages(w http.ResponseWriter, r *http.Request) {
	sId := mux.Vars(r)
	sid := sId["id"]
	id, err := strconv.Atoi(sid)
	fmt.Println(sid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	repository.DeleteTask(id)

}
