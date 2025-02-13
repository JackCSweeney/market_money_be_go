package models

import (
	"testing"
	"example.com/mod/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var Market1 Market
var Market2 Market

func teardownMarkets() {
	database.Db.Exec(`DELETE FROM markets`)
}

func addMarketsToDb() {
	market1 :=  Market{Name: "Market 1", Street: "123 4th St.", City: "Los Angeles", County: "Los Angeles", State: "CA", Zip: "90034", Lat: "123.456", Lon: "78.901"}
	market2 :=  Market{Name: "Market 2", Street: "234 5th St.", City: "Los Angeles", County: "Los Angeles", State: "CA", Zip: "90034", Lat: "123.456", Lon: "78.901"}
	insert1 := `INSERT INTO markets (name, street, city, county, state, zip, lat, lon) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	
	id1 := 0
	id2 := 0
	database.Db.QueryRow(insert1, market1.Name, market1.Street, market1.City, market1.County, market1.State, market1.Zip, market1.Lat, market1.Lon).Scan(&id1)
	database.Db.QueryRow(insert1, market2.Name, market2.Street, market2.City, market2.County, market2.State, market2.Zip, market2.Lat, market2.Lon).Scan(&id2)
	market1.Id = id1
	market2.Id = id2

	Market1 = market1
	Market2 = market2
}

func TestGetMarketByID(t *testing.T) {
	market, _ := GetMarketById(Market1.Id)

	assert.Equal(t, Market1.Id, market.Id)
	assert.Equal(t, Market1.Name, market.Name)
	assert.Equal(t, Market1.Street, market.Street)
	assert.Equal(t, Market1.City, market.City)
	assert.Equal(t, Market1.County, market.County)
	assert.Equal(t, Market1.State, market.State)
	assert.Equal(t, Market1.Zip, market.Zip)
	assert.Equal(t, Market1.Lat, market.Lat)
	assert.Equal(t, Market1.Lon, market.Lon)
}

func TestGetAllMarkets(t *testing.T) {
	markets, _ := GetAllMarkets()

	assert.Equal(t, 2, len(markets))
	assert.Equal(t, Market1.Id, markets[0].Id)
	assert.Equal(t, Market2.Id, markets[1].Id)
}
