package postgres

import (
	"fmt"
	"log"

	"github.com/edelwei88/bytebuild-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB

func OpenPostgresConnection() {
	user := config.Config.Postgres.User
	password := config.Config.Postgres.Password
	dbname := config.Config.Postgres.DB
	port := config.Config.Postgres.Port
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s", user, password, dbname, port)

	var err error
	Postgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
