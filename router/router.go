package router

import (
	"github.com/gin-gonic/gin"
	"example.com/mod/database"
	"example.com/mod/handlers"
	"fmt"
	"strconv"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func Router() {
	router := setupRouter()
	database.ConnectDatabase("DB_NAME", ".env")

	publicRoutes := router.Group("/api")
	// Markets
	publicRoutes.GET("/markets", handlers.HandleGetAllMarkets)
	publicRoutes.GET("/markets/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		handlers.HandleGetMarketById(c, id)
	})
	// // Vendors
	// publicRoutes.GET("/vendors", HandleGetAllVendors())
	// publicRoutes.GET("/vendors/:id", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	// 	HandleGetVendorById()
	// })
	// publicRoutes.POST("/vendors", HandleCreateVendor())
	// publicRoutes.PATCH("/vendors/:id", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	// 	HandleUpdateVendor()
	// })
	// // Market Vendors
	// publicRoutes.GET("/markets/:id/vendors", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	// 	HandleGetAllVendorsForMarket()
	// })
	// publicRoutes.POST("/markets/:id/vendors", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	// 	HandleCreateMarketVendor()
	// })
	// publicRoutes.DELETE("/markets/:id/vendors", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	// 	HandleDeleteMarketVendor()
	// })

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server on port :8080")
		panic(err)
	}

	fmt.Println("Now listening and serving on port :8080")
}
