package main

import (
	"context"
	"paymentweb/internal/configuration"
	"paymentweb/internal/controllers/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	db := configuration.New()
	r := gin.Default()
	routes.Init(&ctx, db, r)
	r.Run()
}
