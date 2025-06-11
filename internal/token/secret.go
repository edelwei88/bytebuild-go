package token

import (
	"github.com/edelwei88/bytebuild-go/internal/config"
)

var JWTSecretKey []byte

func SetupJWT() {
	JWTSecretKey = []byte(config.Config.JWT.Secret)
}
