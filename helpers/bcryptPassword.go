package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) string {
	salt := 8
	password := []byte(p)

	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)

}

func ComparePass(h, p []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
