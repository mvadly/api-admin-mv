package util

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func SaltMD5(username, password string) string {
	salt := os.Getenv("SALT_KEY")
	result := md5.Sum([]byte(username + password + salt))
	final := hex.EncodeToString(result[:])
	return final
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func DefaultPassword(username string) string {
	var defaultPass = username + "P@ssw0rd123"
	pass, _ := HashPassword(defaultPass)
	return pass

}
