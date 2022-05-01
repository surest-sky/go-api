package main

import (
	"github.com/joho/godotenv"
	"go-api/router"
	"os"
)

func main() {
	r := router.Router()
	_ = godotenv.Load(".env", ".env.production")
	_ = r.Run(":" + os.Getenv("APP_PORT"))
}


