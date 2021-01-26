package lib

import "golang.org/x/crypto/bcrypt"

/* Function to hash a plaintext password. */
func Hash(plaintText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintText), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

/* Function to compare plaintext and hash. */
func VerifyHash(hash, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText))
	if err != nil {
		return false
	}
	return true
}
