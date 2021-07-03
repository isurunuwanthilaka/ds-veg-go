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

// Ping ping service
func GetAVegetable(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))

	var vegetable common.Vegetable

	if err := client.Call("Store.Get", id, &vegetable); err != nil {
		fmt.Println("Error:Store.Get()", err)
	} else {
		c.JSON(http.StatusOK,vegetable)
	}
	
}

func init(){
	client, _ = rpc.DialHTTP("tcp", "127.0.0.1:9000") 
}