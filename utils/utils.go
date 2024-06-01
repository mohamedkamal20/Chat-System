package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// HashPasswordMD5 hashes the given password using MD5.
func HashPasswordMD5(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
