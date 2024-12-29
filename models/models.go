package models

import (
	"time"
)

type (

	User struct {
		ID 			uint 	`json:"id"`
		Name 		string 	`json:"name"`
		Password 	string 	`json:"password"`
		Email 		string 	`json:"email"`
	}

	Message struct {
		ID 			uint 		`bson:"_id,omitenpty" json:"id"`
		Content 	string 		`json:"content"`
		From 		uint 		`json:"from"`
		To 			uint 		`json:"to"`
		CreatedAt 	time.Time 	`json:"created_at"`
	}

	Chat struct {
		UUID 			string 		`bson:"_uuid,omitenpty" json:"uuid"`
		MessagesIDs 	[]uint 		`json:"messages"`
		UserIDs			[]uint 		`json:"users"` // uint or user depend of the future needs
		// private / group ?
	}
)