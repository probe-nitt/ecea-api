package utils

import (
	"encoding/base64"

	"github.com/thanhpk/randstr"
)

func GenerateVerificationCode() string {
	code := randstr.String(20)

	verificationCode := Encode(code)
	return verificationCode
}

func Encode(s string) string {
	data := base64.StdEncoding.EncodeToString([]byte(s))
	return data
}

func Decode(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
