package util

import (
	"crypto"

	u "github.com/Truth1984/awadau-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Uuid() string {
	return uuid.New().String()
}

func Sha256(data string) string {
	h := crypto.SHA256.New()
	h.Write([]byte(data))
	return string(h.Sum(nil))
}

func BcryptCompare(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func BcryptHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		EHWarn(err, "BcryptHash - GenerateFromPassword", LogMap(u.Map("password", password), nil))
		return password
	}
	return string(hash)
}
