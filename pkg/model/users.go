package model

type User struct {
	BaseModel
	SocialID    string `json:"social_id"`
	DisplayName string `json:"display_name"`
	ImageUrl    string `json:"image_url"`
	Password    string `json:"password" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
}

func (user *User) TableName() string {
	return "users"
}

type LoginParam struct {
	UserName *string `json:"user_name"`
	Password *string `json:"password"`
}
