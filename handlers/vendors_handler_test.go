package handlers

import (
	"encoding/json"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestGetVendors(t *testing.T) {
	writer := makeRequest("GET", "/api/vendors", nil)
	assert.Equal(t, http.StatusOK, writer.Code)

	var response map[string][]map[string]any
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)

	data := response["data"]
	assert.Equal(t, 2, len(data))

	vendor1 := data[0]
	assert.Equal(t, float64(Vendor1.Id), vendor1["id"])
	assert.Equal(t, Vendor1.ContactName, vendor1["contact_name"])
	assert.Equal(t, Vendor1.ContactPhone, vendor1["contact_phone"])
	assert.Equal(t, Vendor1.Name, vendor1["name"])
	assert.Equal(t, Vendor1.Description, vendor1["description"])
	assert.Equal(t, Vendor1.CreditAccepted, vendor1["credit_accepted"])

	vendor2 := data[1]
	assert.Equal(t, float64(Vendor2.Id), vendor2["id"])
	assert.Equal(t, Vendor2.ContactName, vendor2["contact_name"])
	assert.Equal(t, Vendor2.ContactPhone, vendor2["contact_phone"])
	assert.Equal(t, Vendor2.Name, vendor2["name"])
	assert.Equal(t, Vendor2.Description, vendor2["description"])
	assert.Equal(t, Vendor2.CreditAccepted, vendor2["credit_accepted"])
}

func TestGetVendorById(t *testing.T) {
	url := fmt.Sprintf("/api/vendors/%d", Vendor1.Id)
	writer := makeRequest("GET", url, nil)
	assert.Equal(t, http.StatusOK, writer.Code)

	var response map[string]map[string]any
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)

	vendor1 := response["data"]
	assert.Equal(t, float64(Vendor1.Id), vendor1["id"])
	assert.Equal(t, Vendor1.ContactName, vendor1["contact_name"])
	assert.Equal(t, Vendor1.ContactPhone, vendor1["contact_phone"])
	assert.Equal(t, Vendor1.Name, vendor1["name"])
	assert.Equal(t, Vendor1.Description, vendor1["description"])
	assert.Equal(t, Vendor1.CreditAccepted, vendor1["credit_accepted"])
	
}