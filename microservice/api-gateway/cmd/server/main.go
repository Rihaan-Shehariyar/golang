package main

import (
	"api-gateway/internal/client"
	"api-gateway/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	clients := client.NewClients()

	routes.Setup(r, clients)

	r.Run("8080")

}
