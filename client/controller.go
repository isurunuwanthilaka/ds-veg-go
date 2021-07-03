package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isurunuwanthilaka/ds-veg-go/common"
)

var (
	client *rpc.Client
)

// GetAVegetable returns a vegetable by id
func GetAVegetable(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))

	var vegetable common.Vegetable

	if err := client.Call("Store.Get", id, &vegetable); err != nil {
		fmt.Println("Error:Store.Get()", err)
	} else {
		c.JSON(http.StatusOK,vegetable)
	}
	
}

// GetAllVegetables returns all vegetables
func GetAllVegetables(c *gin.Context) {
	
	var arr []common.Vegetable

	if err := client.Call("Store.GetAll","", &arr); err != nil {
		fmt.Println("Error:Store.GetAll()", err)
	} else {
		c.JSON(http.StatusOK,arr)
	}
	
}

// AddAVegetable add a vegetable
func AddAVegetable(c *gin.Context) {

	var vegetable common.Vegetable

	if err := c.ShouldBindJSON(&vegetable); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}


	if err := client.Call("Store.AddVeg", vegetable, &vegetable); err != nil {
		fmt.Println("Error:Store.AddVeg()", err)
	} else {
		c.JSON(http.StatusOK,vegetable)
	}
	
}

// UpdateAVegetable returns a vegetable by id
func UpdateAVegetable(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))

	var vegetable common.Vegetable

	if err := c.ShouldBindJSON(&vegetable); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	vegetable.ID = id

	if err := client.Call("Store.Update", vegetable, &vegetable); err != nil {
		fmt.Println("Error:Store.Update()", err)
	} else {
		c.JSON(http.StatusOK,vegetable)
	}
	
}

func init(){
	client, _ = rpc.DialHTTP("tcp", "127.0.0.1:9000") 
}