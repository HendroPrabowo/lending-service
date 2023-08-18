package password

import (
	"encoding/base64"
)

func Encrypt(plaintext string) (cipher string) {
	cipher = base64.StdEncoding.EncodeToString([]byte(plaintext))
	return
}

func Decrypt(cipher string) (plaintext string, err error) {
	data, err := base64.StdEncoding.DecodeString(cipher)
	return string(data), err
}
