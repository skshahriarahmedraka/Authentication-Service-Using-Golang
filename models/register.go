package model

type RegModel struct {
	Name string `json:"Name" bson:"Name" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Email    string `json:"Email" bson:"Email" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Password string `json:"Password" bson:"Password" gorm:"type:varchar(100)" validate:"min=2,max=100"`
}
