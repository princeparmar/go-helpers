package utils

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
)

func MD5Hash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

func CreateJWT(secret string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(secret)
}
