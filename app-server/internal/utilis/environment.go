package utilis

import "os"

func getSecretKey() []byte {
	secret_key := os.Getenv("SECRET_KEY")
	secret_byte := []byte(secret_key)
	return secret_byte
}

func GetDatabaseUser() string {
	db_user := os.Getenv("POSTGRES_USER")
	return db_user
}

func GetDatabasePassword() string {
	db_password := os.Getenv("POSTGRES_PASSWORD")
	return db_password
}

func GetDatabaseName() string {
	db_name := os.Getenv("POSTGRES_DB")
	return db_name
}

func GetDatabasePort() string {
	db_port := os.Getenv("POSTGRES_PORT")
	return db_port
}
