package main

import (
	"fmt"
	"strconv"
	"example.com/mod/market_vendors"
	"example.com/mod/database"
	"example.com/mod/markets"
	"example.com/mod/vendors"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	database.ConnectDatabase()

	// vendor endpoints
	route.POST("/vendors", vendors.AddVendor)
	route.PATCH("/vendors", vendors.UpdateVendor)
	route.GET("/vendors", vendors.GetAllVendors)
	route.GET("/vendors/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		vendors.GetOneVendor(c, id)
	})

	// market endpoints
	route.GET("/markets", markets.GetAllMarkets)
	route.GET("/markets/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		markets.GetOneMarket(c, id)
	})

	// market vendor endpoints
	route.POST("/market_vendors", marketvendors.AddMarketVendor)

	err := route.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
}
