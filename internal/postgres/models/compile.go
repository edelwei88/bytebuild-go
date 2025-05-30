package models

type Compile struct {
	ID         uint   `json:"id"`
	CompilerID uint   `json:"compiler_id" gorm:"not null"`
	UserID     uint   `json:"user_id" gorm:"not null"`
	User       User   `json:"user"`
	Arg        string `json:"args"`
	Stdout     string `json:"stdout"`
	Stderr     string `json:"stderr"`
}
