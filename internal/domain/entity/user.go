package entity

type User struct {
	UserId   string `gorm:"primaryKey;" json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
