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

func InsertTask(msg MessageRequest) Message {
	message := Message{
		Task:   msg.Message,
		IsDone: false,
	}
	DB.Create(&message)
	log.Println("Successfully added")
	return message
}

func UpdateTask(iD int, msg MessageRequest) Message {
	var message Message
	DB.Model(&message).Where("id = ?", iD).Updates(Message{Task: msg.Message})
	DB.First(&message, iD)
	return message
}

type MessageRequest struct {
	Message string `json:"message"`
}

func DeleteTask(iD int) {
	DB.Delete(&Message{}, iD)
}

func GetMessages(messages *[]Message) {
	DB.Find(&messages)
}
