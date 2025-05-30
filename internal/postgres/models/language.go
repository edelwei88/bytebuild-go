package models

type Language struct {
	ID            uint       `json:"id"`
	Name          string     `json:"name" gorm:"not null"`
	FileExtension string     `json:"file_extension" gorm:"not null"`
	Compilers     []Compiler `json:"compilers" gorm:"constraint:OnDelete:CASCADE"`
}
