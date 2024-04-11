package jwts

import "github.com/dgrijalva/jwt-go"

// JwtPayLoad jwt负载
type JwtPayLoad struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Role     int    `json:"role"`
}

var JwtSecretKey []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
