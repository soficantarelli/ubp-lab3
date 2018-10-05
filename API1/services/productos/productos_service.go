package productos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	productosDomain "github.com/soficantarelli/ubp-lab3/API1/domain/productos"
)

var (
	productList = make([]productosDomain.Producto, 0)
)

func LoadProducts() {
	bytes, err := ioutil.ReadFile("./productos.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(bytes, &productList); err != nil {
		panic(err)
	}

	fmt.Println(productList)
}

func GetProductList() *[]productosDomain.Producto {
	return &productList
}

func AddProduct(body []byte) {

	var p productosDomain.Producto
	json.Unmarshal(body, &p)
	productList = append(productList, p)

	bytes, _ := json.Marshal(&productList)

	ioutil.WriteFile("./productos.json", bytes, 0644)

}
