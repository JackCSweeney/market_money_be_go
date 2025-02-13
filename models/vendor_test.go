package models

import (
	"testing"
	"example.com/mod/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var Vendor1 Vendor
var Vendor2 Vendor

func teardownVendors() {
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
	assert.Equal(t, Vendor1.Id, vendors[0].Id)
	assert.Equal(t, Vendor2.Id, vendors[1].Id)
}

func TestUpdateVendor(t *testing.T) {
	vendorUpdates := Vendor{Id: Vendor1.Id, Name: "Vendor 3", Description: "Third Vendor", ContactName: "Updated Name", ContactPhone: "New Num", CreditAccepted: true}
	vendor, _ := UpdateVendor(vendorUpdates.Id, vendorUpdates)
	var updatedVendor Vendor
	updatedVendor, _ = GetVendorById(vendor.Id)

	assert.Equal(t, "Updated Name", vendor.ContactName)
	assert.Equal(t, "Updated Name", updatedVendor.ContactName)
}

func TestCreateVendor(t *testing.T) {
	newVendor, _ := CreateVendor("Name", "Description", "Contact Name", "Contact Phone", false)
	foundVendor, _ := GetVendorById(newVendor.Id)
	assert.IsType(t, Vendor{}, newVendor)
	assert.Equal(t, foundVendor.Id, newVendor.Id)
	assert.Equal(t, foundVendor.Name, newVendor.Name)
	assert.Equal(t, foundVendor.ContactName, newVendor.ContactName)
	assert.Equal(t, foundVendor.ContactPhone, newVendor.ContactPhone)
	assert.Equal(t, foundVendor.Description, newVendor.Description)
	assert.Equal(t, foundVendor.CreditAccepted, newVendor.CreditAccepted)
}