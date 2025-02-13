package models

import (
	"testing"
	"example.com/mod/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var MarketVendor1 MarketVendor
var MarketVendor2 MarketVendor
var MarketVendor3 MarketVendor

func teardownMarketVendors() {
	database.Db.Exec(`DELETE FROM market_vendors`)
}

func addMarketVendorsToDb() {
	MarketVendor1 = MarketVendor{MarketId: Market1.Id, VendorId: Vendor1.Id}
	MarketVendor2 = MarketVendor{MarketId: Market1.Id, VendorId: Vendor2.Id}
	MarketVendor3 = MarketVendor{MarketId: Market2.Id, VendorId: Vendor1.Id}

	insert := `INSERT INTO market_vendors (market_id, vendor_id) VALUES ($1, $2) RETURNING id`
	database.Db.QueryRow(insert, MarketVendor1.MarketId, MarketVendor1.VendorId).Scan(&MarketVendor1.Id)
	database.Db.QueryRow(insert, MarketVendor2.MarketId, MarketVendor2.VendorId).Scan(&MarketVendor2.Id)
	database.Db.QueryRow(insert, MarketVendor3.MarketId, MarketVendor3.VendorId).Scan(&MarketVendor3.Id)
}

func TestGetVendorsForMarket(t *testing.T) {
	vendors1, _ := GetVendorsForMarket(Market1.Id)
	vendors2, _ := GetVendorsForMarket(Market2.Id)

	assert.Equal(t, 2, len(vendors1))
	assert.Equal(t, 1, len(vendors2))
	assert.Equal(t, Vendor1.Id, vendors1[0].Id)
	assert.Equal(t, Vendor2.Id, vendors1[1].Id)
	assert.Equal(t, Vendor1.Id, vendors2[0].Id)
}

// func CreateMarketVendor(t *testing.T) {

// }

// func DeleteMarketVendor(t *testing.T) {
	
// }

