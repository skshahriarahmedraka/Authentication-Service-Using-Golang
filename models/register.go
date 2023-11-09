package model

type RegModel struct {
	Firstname string `json:"firstname" bson:"firstname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Lastname  string `json:"lastname" bson:"lastname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Email    string `json:"Email" bson:"Email" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Password string `json:"Password" bson:"Password" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	
}
