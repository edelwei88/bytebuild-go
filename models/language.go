package models

type Language struct {
	ID         uint
	Name       string     `gorm:"not null"`
	Extenstion string     `gorm:"not null"`
	Compilers  []Compiler `gorm:"constraint:OnDelete:CASCADE"`
}
