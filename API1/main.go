package main

import (
	"github.com/gin-gonic/gin"
	productosController "github.com/soficantarelli/ubp-lab3/API1/controllers/productos"
	productosService "github.com/soficantarelli/ubp-lab3/API1/services/productos"
)

func main() {
	productosService.LoadProducts()

	router := gin.Default()

	router.GET("/productos", productosController.GetProductos)
	router.POST("/productos", productosController.PostProductos)

	router.Run(":8080")
}
