package token

import (
	"log"
	"os"
)

var JWTSecretKey []byte

func InitJWT() {
	JWTSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(JWTSecretKey) == 0 {
		log.Fatal("failed to init JWT secret key")
	}
}
