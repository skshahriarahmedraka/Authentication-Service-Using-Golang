package model

type LoginModel struct {
	//UserName string `json:"UserName" bson:"UserName"`
	Email    string `json:"email" bson:"email" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Password string `json:"password" bson:"password" gorm:"type:varchar(100)" validate:"min=2,max=100"`
}
