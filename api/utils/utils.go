package utils

import (
	env "github.com/joho/godotenv"
)

// LoadEnv should load .env file
func LoadEnv(){
	env.Load()
	///Users/kimberly.luna/go/src/github.com/kimberly.luna/proxy-app/.env
	// fmt.Println(os.Getenv("PORT"))
}
