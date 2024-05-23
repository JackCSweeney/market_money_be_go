package main

import (
	"fmt"
	"example.com/mod/vendors"
	"example.com/mod/database"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	database.ConnectDatabase()
	route.POST("/vendors", vendors.AddVendor)

	err := route.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
}