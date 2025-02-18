package models

import (
	"testing"
	"example.com/mod/database"
	"os"
	_ "github.com/lib/pq"
)


func setup() {
	database.ConnectDatabase("TEST_DB_NAME", "../.env")
}

func TestMain(m *testing.M) {
	setup()
	addVendorsToDb()
	addMarketsToDb()
	addMarketVendorsToDb()
	exitCode := m.Run()
	teardownMarketVendors()
	teardownVendors()
	teardownMarkets()
	os.Exit(exitCode)
}