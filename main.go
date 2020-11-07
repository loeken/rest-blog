package main

import (
	"fmt"

	"github.com/loeken/rest-blog/api"
)

// @title Rest Blog
// @version 1.0
// @description This is a sample service for a blog
// @termsOfService loekenIsKing
// @contact.name loeken
// @contact.email loeken@internetz.me
// @host localhost:8080
// @BasePath /api/v1
// @Schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @tokenUrl http://localhost:8080/api/v1/auth
func main() {
	fmt.Println("application started")
	api.Run()
}
