package main

import (
	"github.com/books/models"
	"github.com/books/routes"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	// r := gin.Default()
	godotenv.Load()          // Load env variables
	models.ConnectDataBase() // load db

	var router = make(chan *gin.Engine)
	go routes.GetRouter(router)
	r := <-router
	r.Run()
}
