package models

type User struct {
	ID       uint      `json:"id"`
	Username string    `json:"username" gorm:"not null"`
	Email    string    `json:"email" gorm:"not null"`
	Password string    `json:"-" gorm:"not null"`
	RoleID   string    `json:"role_id" gorm:"not null"`
	Role     Role      `json:"role"`
	Compiles []Compile `json:"compiles" gorm:"constraint:OnDelete:CASCADE"`
}
