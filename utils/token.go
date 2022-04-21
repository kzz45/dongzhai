package utils

import (
	"dongzhai/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   uint   `json:"role_id"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

func CreateToken(id uint, username, password string, role_id uint, admin bool) (string, error) {
	base_claims := Claims{
		id,
		username,
		password,
		role_id,
		admin,
		jwt.StandardClaims{
			Issuer:    "issuer",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, base_claims)
	token, err := tokenClaims.SignedString([]byte(config.GlobalConfig.Server.Secret))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GlobalConfig.Server.Secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
