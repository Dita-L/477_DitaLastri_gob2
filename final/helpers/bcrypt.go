package helpers

import "golang.org/x/crypto/bcrypt"

//membuat hash
func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	strHash := string(hash)
	return strHash
}

//mengecek password
func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
