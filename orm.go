package main

import (
	"log"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}

func InsertTask(msg MessageRequest) {
	message := Message{
		Task: msg.Message, IsDone: false,
	}
	DB.Create(&message)
	log.Println("Succesfully added")
}
