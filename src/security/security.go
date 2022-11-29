package security

import "golang.org/x/crypto/bcrypt"

// Hash generates a hash for the given password
func Hash(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
}

// VerifyPW compares if the given hash matches the password
func VerifyPW(hashedpw, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpw), []byte(pw))
}
