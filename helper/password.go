package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), err
}

func CheckPassword(raw, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))
}