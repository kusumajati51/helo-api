package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) string {

	pass := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	return string(hash)
}

func ComparePassword(userPassword string, inputPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(inputPassword))
	return err != nil
}
