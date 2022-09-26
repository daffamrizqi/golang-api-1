package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// routers
	router.GET("/", rootHandler)
	router.GET("/car/:id", carHandler)
	router.GET("/query", queryHandler)

	router.POST("car", carInput)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":  "DAuto-Luxury-Garage",
		"motto": "Made by enthusiast for enthusiasts",
	})
}

func carHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"car": id,
	})
}

func queryHandler(c *gin.Context) {
	brand := c.Query("brand")
	c.JSON(http.StatusOK, gin.H{
		"brand": brand,
	})
}

type CarInput struct {
	Brand string
	Price int
}

func carInput(c *gin.Context) {
	var carInput CarInput

	err := c.ShouldBindJSON(&carInput)

	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"brand": carInput.Brand,
		"price": carInput.Price,
	})
}
