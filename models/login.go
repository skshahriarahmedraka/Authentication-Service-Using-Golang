package model

type LoginModel struct {
	//UserName string `json:"UserName" bson:"UserName"`
	Email    string `json:"Email" bson:"Email" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Password string `json:"Password" bson:"Password" gorm:"type:varchar(100)" validate:"min=2,max=100"`
}
