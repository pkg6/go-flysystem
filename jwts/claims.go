package jwts

import "github.com/golang-jwt/jwt"

type FlysystemClaims struct {
	jwt.StandardClaims
	Iss    string `json:"iss"`
	Iat    int64  `json:"iat"`
	Exp    int64  `json:"exp"`
	Aud    string `json:"aud"`
	Disk   string `json:"disk"`
	Bucket string `json:"bucket"`
}
