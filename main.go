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

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func main() {
	router := setupRouter()
	database.ConnectDatabase()

	// vendor endpoints
	router.POST("/vendors", vendors.AddVendor)
	router.PATCH("/vendors", vendors.UpdateVendor)
	router.GET("/vendors", vendors.GetAllVendors)
	router.GET("/vendors/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		vendors.GetOneVendor(c, id)
	})

	// market endpoints
	router.GET("/markets", markets.GetAllMarkets)
	router.GET("/markets/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		markets.GetOneMarket(c, id)
	})

	// market vendor endpoints
	router.POST("/market_vendors", marketvendors.AddMarketVendor)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
}
