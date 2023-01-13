package main

import (
	"fmt"
	" hery-ciaputra/demo-gin/handlers"
	" hery-ciaputra/demo-gin/models"
	"github.com/gin-gonic/gin"
)

var products = make(map[int]*models.Product)

func myMiddleware(text string) gin.HandlerFunc {
	return func(c *gin.Context) {
		name, _ := c.Get("User")
		fmt.Println("middleware executed", text, name)
		if text == "validateMdw" {
			fmt.Println("abort request")
			c.AbortWithStatusJSON(500, gin.H{
				"message": "validate failed",
			})
		}
		c.Next()
	}
}

func main() {
	err := db.Connect()
	if err != nil {
		fmt.Println("failed to connect to db")
	}
	//router := gin.Default()
	//
	//products[1] = &models.Product{
	//	ID:          1,
	//	Name:        "Macbook Pro",
	//	Description: "Macbook yang keren",
	//	Quantity:    12,
	//}

	// */products/:id => request resource with specific ID
	// if not found => 404

	// /products => [product]
	// */products?name=acer => request resource with specific specific_ID
	// if not found still 200 => []

	productRepository := repositories.NewProductRepository(&repositories)

	h := handlers.Handler{}

	router.Use(myMiddleware("Global middleware"))
	router.GET("/products/:id", myMiddleware("before handlers"), handler, myMiddleware("after handlers"))

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
