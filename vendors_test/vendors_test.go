package vendors_test

import (
	"testing"
	"net/http/httptest"
	"os"
	"example.com/mod/vendors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
)

var Router *gin.Engine

func TestMain(m *testing.M){
	SetupTestDB()
	addVendorsToDb()
	code := m.Run()
	TearDownTestDB()
	os.Exit(code)
}

func addVendorsToDb() {
	vendor1 := vendors.Vendor{Name: "Vendor 1", Description: "First Vendor", ContactName: "Vendor Name 1", ContactPhone: "18002329393", CreditAccepted: false}
	vendor2 := vendors.Vendor{Name: "Vendor 2", Description: "Second Vendor", ContactName: "Vendor Name 2", ContactPhone: "18002329393", CreditAccepted: true}

	insert1 := fmt.Sprintf("INSERT INTO vendors VALUES %s, %s, %s, %s", vendor1.Name, vendor1.Description, vendor1.ContactName, vendor1.ContactPhone)
	insert2 := fmt.Sprintf("INSERT INTO vendors VALUES %s, %s, %s, %s", vendor2.Name, vendor2.Description, vendor2.ContactName, vendor2.ContactPhone)

	Db.Query(insert1)
	Db.Query(insert2)
}

func TestGetAllVendors(t *testing.T) {
	router := gin.Default()
	router.GET("/vendors", vendors.GetAllVendors)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/vendors", nil)

	router.ServeHTTP(w, req)
	resp, _ := io.ReadAll(w.Body)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(resp))
}