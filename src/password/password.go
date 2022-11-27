package password

import "golang.org/x/crypto/bcrypt"

func Hash(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
}

func VerifyPW(hashedpw, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpw), []byte(pw))
}
