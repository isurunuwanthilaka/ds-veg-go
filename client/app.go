package main

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main(){
	// URL mapping
	router.GET("/vegetable/:id", GetAVegetable)

	router.Run(":8080")
}