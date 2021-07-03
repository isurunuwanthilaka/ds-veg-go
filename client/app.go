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
	router.GET("/vegetable", GetAllVegetables)
	router.POST("/vegetable", AddAVegetable)
	router.POST("/vegetable/:id", UpdateAVegetable)
	router.DELETE("/vegetable/:id", DeleteAVegetable)

	router.Run(":8080")
}