package pkg

import "golang.org/x/crypto/bcrypt"

// encrypt password
func Encrypt(source string) (string, error) {
	hasPwd, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hasPwd), err
}

// compare password
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
}
