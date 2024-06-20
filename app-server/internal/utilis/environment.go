package utilis

import "os"

func getSecretKey() []byte {
	secret_key := os.Getenv("SECRET_KEY")
	secret_byte := []byte(secret_key)
	return secret_byte
}
