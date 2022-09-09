package utils

import "golang.org/x/crypto/bcrypt"

func Hash(pwd string) (string, error) {
	pwdCrypt, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(pwdCrypt), err
}
