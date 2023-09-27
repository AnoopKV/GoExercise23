package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func BcryptString(str string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Println("Hashed::" + string(res))
	return string(res), nil
}

// return nil if matches, or else error will be thrown
func ComparePassword(hashedPwd string, pwd string) error {
	return bcrypt.CompareHashAndPassword(([]byte(hashedPwd)), []byte(pwd))
}
