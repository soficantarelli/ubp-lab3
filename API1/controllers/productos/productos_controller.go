package productos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	productosService "github.com/soficantarelli/ubp-lab3/API1/services/productos"
	"github.com/soficantarelli/ubp-lab3/API1/utils"
)

func GetProductos(c *gin.Context) {

	productList := productosService.GetProductList()

	c.JSON(http.StatusOK, productList)
}

func PostProductos(c *gin.Context) {

	body := utils.GetJsonBody(c.Request)

	productosService.AddProduct(body)

	c.String(http.StatusCreated, "ok")
}
