// package main

// import (
// 	"testing"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"fmt"
// 	"github.com/gin-gonic/gin"
// 	"database/sql"
// 	"github.com/joho/godotenv"
// 	"strconv"
// 	"example.com/mod/models"
// 	"encoding/json"
// 	"bytes"
// )

// func TestMain(m *testing.M) {
// 	gin.SetMode(gin.TestMode)
// 	setupTestDB()
// 	exitCode := m.Run()
// 	tearDownTestDB()

// 	os.Exit(exitCode)
// }

// func router() *gin.Engine {
// 	router := gin.Default()

// 	// vendor endpoints
// 	router.POST("/vendors", vendors.AddVendor)
// 	router.PATCH("/vendors", vendors.UpdateVendor)
// 	router.GET("/vendors", vendors.GetAllVendors)
// 	router.GET("/vendors/:id", func(c *gin.Context) {
// 		id, _ := strconv.Atoi(c.Param("id"))
// 		vendors.GetOneVendor(c, id)
// 	})

// 	return router
// }

// func MakeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
// 	requestBody, _ := json.Marshal(body)
// 	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
// 	writer := httptest.NewRecorder()
// 	router().ServeHTTP(writer, request)
// 	return writer
// }

// var Db *sql.DB

// func setupTestDB(){
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	host := os.Getenv("HOST")
// 	port, _ := strconv.Atoi(os.Getenv("PORT"))
// 	user := os.Getenv("USER")
// 	dbname := "market_money_test"

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"dbname=%s sslmode=disable", host, port, user, dbname)

// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("There was an error while connection to the database:", err)
// 		panic(err)
// 	} else {
// 		Db = db
// 		fmt.Println("Successfully connected to database")
// 	}
// }

// func tearDownTestDB() {
// 	Db.Query("DELETE * FROM vendors, markets, market_vendors")
// }