package models

type Compiler struct {
	ID              uint      `json:"id"`
	DockerImageName string    `json:"docker_image_name" gorm:"not null"`
	LanguageID      uint      `json:"language_id" gorm:"not null"`
	Compiles        []Compile `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}
