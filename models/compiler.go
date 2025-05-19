package models

type Compiler struct {
	ID              uint
	DockerImageName string `gorm:"not null"`
	LanguageID      uint   `gorm:"not null" json:"-"`
}
