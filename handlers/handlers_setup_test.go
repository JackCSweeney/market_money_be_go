package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"example.com/mod/database"
	"example.com/mod/models"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}

func router() *gin.Engine {
	router := gin.Default()
	
	publicRoutes := router.Group("/api")
	// Markets
	publicRoutes.GET("/markets", HandleGetAllMarkets)
	publicRoutes.GET("/markets/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		HandleGetMarketById(c, id)
	})
	
	// Vendors
	publicRoutes.GET("/vendors", HandleGetAllVendors)
	publicRoutes.GET("/vendors/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		HandleGetVendorById(c, id)
	})
	// publicRoutes.POST("/vendors", HandleCreateVendor)
	publicRoutes.PATCH("/vendors/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		HandleUpdateVendor(c, id)
	})
	// // Market Vendors
	// publicRoutes.GET("/markets/:id/vendors", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	// 	HandleGetAllVendorsForMarket(c, id)
	// })
	// publicRoutes.POST("/markets/:id/vendors", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	// 	HandleCreateMarketVendor(c, id)
	// })
	// publicRoutes.DELETE("/markets/:id/vendors", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	// 	HandleDeleteMarketVendor(c, id)
	// })

	return router
}

func setup() {
	database.ConnectDatabase("TEST_DB_NAME", "../.env")
	addMarketsToDb()
	addVendorsToDb()
	addMarketVendorsToDb()
}

func teardown() {
	teardownMarketVendors()
	teardownMarkets()
	teardownVendors()
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer 
}

// Models and DB Setup that can likely be moved to a separate file
var Market1 models.Market
var Market2 models.Market

func teardownMarkets() {
	database.Db.Exec(`DELETE FROM markets`)
}

func addMarketsToDb() {
	market1 :=  models.Market{Name: "Market 1", Street: "123 4th St.", City: "Los Angeles", County: "Los Angeles", State: "CA", Zip: "90034", Lat: "123.456", Lon: "78.901"}
	market2 :=  models.Market{Name: "Market 2", Street: "234 5th St.", City: "Los Angeles", County: "Los Angeles", State: "CA", Zip: "90034", Lat: "123.456", Lon: "78.901"}
	insert := `INSERT INTO markets (name, street, city, county, state, zip, lat, lon) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	
	id1 := 0
	id2 := 0
	database.Db.QueryRow(insert, market1.Name, market1.Street, market1.City, market1.County, market1.State, market1.Zip, market1.Lat, market1.Lon).Scan(&id1)
	database.Db.QueryRow(insert, market2.Name, market2.Street, market2.City, market2.County, market2.State, market2.Zip, market2.Lat, market2.Lon).Scan(&id2)
	market1.Id = id1
	market2.Id = id2

	Market1 = market1
	Market2 = market2
}

var Vendor1 models.Vendor
var Vendor2 models.Vendor

func teardownVendors() {
	database.Db.Exec(`DELETE FROM vendors`)
}

func addVendorsToDb() {
	vendor1 := models.Vendor{Name: "Vendor 1", Description: "First Vendor", ContactName: "Vendor Name 1", ContactPhone: "18002329393", CreditAccepted: false}
	vendor2 := models.Vendor{Name: "Vendor 2", Description: "Second Vendor", ContactName: "Vendor Name 2", ContactPhone: "18002329393", CreditAccepted: true}

	insert := `INSERT INTO vendors (name, description, contact_name, contact_phone, credit_accepted) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	
	id1 := 0
	database.Db.QueryRow(insert, vendor1.Name, vendor1.Description, vendor1.ContactName, vendor1.ContactPhone, vendor1.CreditAccepted).Scan(&id1)
	vendor1.Id = id1

	id2 := 0
	database.Db.QueryRow(insert, vendor2.Name, vendor2.Description, vendor2.ContactName, vendor2.ContactPhone, vendor2.CreditAccepted).Scan(&id2)
	vendor2.Id = id2

	Vendor1 = vendor1
	Vendor2 = vendor2
}

var MarketVendor1 models.MarketVendor
var MarketVendor2 models.MarketVendor
var MarketVendor3 models.MarketVendor

func teardownMarketVendors() {
	database.Db.Exec(`DELETE FROM market_vendors`)
}

func addMarketVendorsToDb() {
	MarketVendor1 = models.MarketVendor{MarketId: Market1.Id, VendorId: Vendor1.Id}
	MarketVendor2 = models.MarketVendor{MarketId: Market1.Id, VendorId: Vendor2.Id}
	MarketVendor3 = models.MarketVendor{MarketId: Market2.Id, VendorId: Vendor1.Id}

	insert := `INSERT INTO market_vendors (market_id, vendor_id) VALUES ($1, $2) RETURNING id`
	database.Db.QueryRow(insert, MarketVendor1.MarketId, MarketVendor1.VendorId).Scan(&MarketVendor1.Id)
	database.Db.QueryRow(insert, MarketVendor2.MarketId, MarketVendor2.VendorId).Scan(&MarketVendor2.Id)
	database.Db.QueryRow(insert, MarketVendor3.MarketId, MarketVendor3.VendorId).Scan(&MarketVendor3.Id)
}