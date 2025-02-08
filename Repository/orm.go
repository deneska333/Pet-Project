package repository

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

func UpdateTask(iD int, msg MessageRequest) {
	var message Message
	DB.Model(&message).Where("id = ?", iD).Updates(Message{Task: msg.Message})
}

type MessageRequest struct {
	Message string `json:"message"`
}

func DeleteTask(iD int) {
	DB.Delete(&Message{}, iD)
}
