package tool

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
)

// ValidToken token jwt.Token return string
func ValidToken(t *jwt.Token) string {
	claims := t.Claims.(jwt.MapClaims)
	return claims["name"].(string)
}

// PassMd5 pass string return string
func PassMd5(pass string) string {
	str := md5.New()
	str.Write([]byte(pass))
	return hex.EncodeToString(str.Sum(nil))
}