package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type UserData struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Firstname string `json:"firstname" bson:"firstname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Lastname  string `json:"lastname" bson:"lastname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Email     string `json:"email" bson:"email" gorm:"unique;not null;type:varchar(100)" validate:"required,min=3,max=100"`
	Password  string `json:"password" bson:"password" gorm:"type:varchar(100)" validate:"required,min=3,max=50"`
	Telephone string `json:"telephone" bson:"telephone" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Address string `json:"address" bson:"address" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	AccountType string `json:"accounttype" bson:"accounttype" gorm:"type:varchar(100)" validate:"min=2,max=100"`
}