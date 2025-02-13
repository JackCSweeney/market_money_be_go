package models

import (
	"testing"
	"example.com/mod/database"
	"os"
	_ "github.com/lib/pq"
)


func setup() {
	database.ConnectDatabase("TEST_DB_NAME")
}

func TestMain(m *testing.M) {
	setup()
	addVendorsToDb()
	addMarketsToDb()
	exitCode := m.Run()
	teardownVendors()
	teardownMarkets()
	os.Exit(exitCode)
}