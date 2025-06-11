package models

type Compile struct {
	ID          uint     `json:"id"`
	CompilerID  uint     `json:"compiler_id" gorm:"not null"`
	Compiler    Compiler `json:"compiler"`
	UserID      uint     `json:"user_id" gorm:"not null"`
	User        User     `json:"-"`
	SourceCode  string   `json:"source_code" gorm:"not null"`
	Arg         string   `json:"args"`
	ExitCode    int      `json:"exit_code" gorm:"not null"`
	Stdout      string   `json:"stdout"`
	Stderr      string   `json:"stderr"`
	CompileTime int64    `json:"compile_time" gorm:"not null"`
}
