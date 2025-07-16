package routes

import (
	"crud-golang/backend/controllers"

	"github.com/gin-gonic/gin"
)


func ProductRoutes(router *gin.Engine) {
	//inisiasi mongo db (panggil sekali saat start)
	controllers.InitMongo()

	// Endpoint GET untuk ambil semua produk
	router.GET("/products", controllers.GetProducts)

	// Endpoint POST untuk tambah produk
	router.POST("/products", controllers.CreateProduct)

	// Endpoint PUT untuk update produk (by id)
	router.PUT("/products/:id", controllers.UpdateProduct)

	// Endpoint DELETE untuk hapus produk (by id)
	router.DELETE("/products/:id", controllers.DeleteProduct)
	
	
}