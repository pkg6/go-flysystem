package token

import "github.com/golang-jwt/jwt"

type FlysystemClaims struct {
	jwt.StandardClaims
	Disk   string `json:"disk"`
	Bucket string `json:"bucket"`
}
