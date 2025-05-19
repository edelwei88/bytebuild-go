package models

type Compile struct {
	ID         uint
	CompilerID uint `gorm:"not null"`
	Compiler   Compiler
	Stdin      string
	Stdout     string
	Stderr     string
}
