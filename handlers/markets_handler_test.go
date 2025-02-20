package handlers

import (
	"encoding/json"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestGetMarkets(t *testing.T) {
	writer := makeRequest("GET", "/api/markets", nil)
	assert.Equal(t, http.StatusOK, writer.Code)

	var response map[string][]map[string]any
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]	
	assert.Equal(t, true, exists)

	data := response["data"]
	assert.Equal(t, 2, len(data))

	market1 := data[0]
	assert.Equal(t, float64(Market1.Id), market1["id"])
	assert.Equal(t, Market1.Name, market1["name"])
	assert.Equal(t, Market1.Street, market1["street"])
	assert.Equal(t, Market1.City, market1["city"])
	assert.Equal(t, Market1.County, market1["county"])
	assert.Equal(t, Market1.State, market1["state"])
	assert.Equal(t, Market1.Zip, market1["zip"])
	assert.Equal(t, Market1.Lat, market1["lat"])
	assert.Equal(t, Market1.Lon, market1["lon"])

	market2 := data[1]
	assert.Equal(t, float64(Market2.Id), market2["id"])
	assert.Equal(t, Market2.Name, market2["name"])
	assert.Equal(t, Market2.Street, market2["street"])
	assert.Equal(t, Market2.City, market2["city"])
	assert.Equal(t, Market2.County, market2["county"])
	assert.Equal(t, Market2.State, market2["state"])
	assert.Equal(t, Market2.Zip, market2["zip"])
	assert.Equal(t, Market2.Lat, market2["lat"])
	assert.Equal(t, Market2.Lon, market2["lon"])
}

func TestGetOneMarket(t *testing.T) {
	url := fmt.Sprintf("/api/markets/%d", Market1.Id)

	writer := makeRequest("GET", url, nil)
	assert.Equal(t, http.StatusOK, writer.Code)

	var response map[string]map[string]any
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)

	market1 := response["data"]
	assert.Equal(t, float64(Market1.Id), market1["id"])
	assert.Equal(t, Market1.Name, market1["name"])
	assert.Equal(t, Market1.Street, market1["street"])
	assert.Equal(t, Market1.City, market1["city"])
	assert.Equal(t, Market1.County, market1["county"])
	assert.Equal(t, Market1.State, market1["state"])
	assert.Equal(t, Market1.Zip, market1["zip"])
	assert.Equal(t, Market1.Lat, market1["lat"])
	assert.Equal(t, Market1.Lon, market1["lon"])
}