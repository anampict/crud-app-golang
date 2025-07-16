package main

import (
	"crud-golang/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.ProductRoutes(router)

	router.Run(":8080")
}
