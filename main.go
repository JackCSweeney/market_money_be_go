package main

import (
	"fmt"
	"example.com/mod/vendors"
	"example.com/mod/database"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	route := gin.Default()
	database.ConnectDatabase()
	route.POST("/vendors", vendors.AddVendor)
	route.PATCH("/vendors", vendors.UpdateVendor)
	route.GET("/vendors", vendors.GetAllVendors)
	route.GET("/vendors/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		vendors.GetOneVendor(c, id)
	})

	err := route.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
}