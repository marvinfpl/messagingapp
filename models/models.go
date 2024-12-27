package models

import (
	"time"
)

type (

	User struct {
		ID 			uint 	`json:"id" gorm:"primaryKey"`
		Name 		string 	`json:"name" gorm:"not null"`
		Password 	string 	`json:"password" gorm:"not null"`
		Email 		string 	`json:"email" gorm:"unique; not null"`
	}

	Message struct {
		ID 			uint 		`bson:"_id,omitenpty" json:"id"`
		Content 	string 		`json:"content"`
		From 		uint 		`json:"from"`
		To 			uint 		`json:"to"`
		CreatedAt 	time.Time 	`json:"created_at"`
	}

	ChatRoom struct {
		UUID 		uint 		`bson:"_uuid,omitenpty" json:"uuid"`
		Messages 	[]Message 	`json:"messages"`
		Users 		[]User 		`json:"users"`
		// private / group ?
	}
)