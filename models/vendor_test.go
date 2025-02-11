package models

import (
	"testing"
	"example.com/mod/database"
	"os"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var Vendor1 Vendor
var Vendor2 Vendor

func setup() {
	database.ConnectDatabase("TEST_DB_NAME")
}

func teardown() {
	database.Db.Exec(`DELETE FROM vendors`)
}

func addVendorsToDb() {
	vendor1 := Vendor{Name: "Vendor 1", Description: "First Vendor", ContactName: "Vendor Name 1", ContactPhone: "18002329393", CreditAccepted: false}
	vendor2 := Vendor{Name: "Vendor 2", Description: "Second Vendor", ContactName: "Vendor Name 2", ContactPhone: "18002329393", CreditAccepted: true}

	insert1 := `INSERT INTO vendors (name, description, contact_name, contact_phone, credit_accepted) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	insert2 := `INSERT INTO vendors (name, description, contact_name, contact_phone, credit_accepted) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	
	id1 := 0
	database.Db.QueryRow(insert1, vendor1.Name, vendor1.Description, vendor1.ContactName, vendor1.ContactPhone, vendor1.CreditAccepted).Scan(&id1)
	vendor1.Id = id1

	id2 := 0
	database.Db.QueryRow(insert2, vendor2.Name, vendor2.Description, vendor2.ContactName, vendor2.ContactPhone, vendor2.CreditAccepted).Scan(&id2)
	vendor2.Id = id2

	Vendor1 = vendor1
	Vendor2 = vendor2
}

func TestMain(m *testing.M) {
	setup()
	addVendorsToDb()
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}

func TestGetVendorByID(t *testing.T) {
	vendor, _ := GetVendorById(Vendor1.Id)

	assert.Equal(t, Vendor1.Id, vendor.Id)
	assert.Equal(t, Vendor1.Name, vendor.Name)
	assert.Equal(t, Vendor1.Description, vendor.Description)
	assert.Equal(t, Vendor1.ContactName, vendor.ContactName)
	assert.Equal(t, Vendor1.ContactPhone, vendor.ContactPhone)
	assert.Equal(t, Vendor1.CreditAccepted, vendor.CreditAccepted)
}

func TestGetAllVendors(t *testing.T) {
	vendors, _ := GetAllVendors()

	assert.Equal(t, 2, len(vendors))
}
