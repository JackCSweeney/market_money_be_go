package models

import (
	"example.com/mod/database"
	"fmt"
)

type MarketVendor struct {
	Id int `json:"id"`
	MarketId int `json:"market_id"`
	VendorId int `json:"vendor_id"`
}

func GetVendorsForMarket(id int) ([]Vendor, error) {
	sqlStatement := `SELECT vendors.* FROM vendors JOIN market_vendors ON vendors.id = market_vendors.vendor_id WHERE market_vendors.market_id = $1`
	var vendors []Vendor

	rows, err := database.Db.Query(sqlStatement, id)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var vendor Vendor
		err = rows.Scan(&vendor.Id, &vendor.Name, &vendor.Description, &vendor.ContactName, &vendor.ContactPhone, &vendor.CreditAccepted)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		vendors = append(vendors, vendor)
	}
	err = rows.Err()
	if err != nil {
		return vendors, err
	}
	return vendors, nil
}

func CreateMarketVendor(marketId int, vendorId int) (MarketVendor, error) {
	sqlStatement := `INSERT INTO market_vendors (market_id, vendor_id) VALUES ($1, $2) RETURNING *`
	var marketVendor MarketVendor

	err := database.Db.QueryRow(sqlStatement, marketId, vendorId).Scan(&marketVendor.Id, &marketVendor.MarketId, &marketVendor.VendorId)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return marketVendor, err
	}
	return marketVendor, nil
}

func DeleteMarketVendor(marketId int, vendorId int) (string, error) {
	sqlStatement := `DELETE FROM market_vendors WHERE market_id=$1 AND vendor_id=$2`
	_, err := database.Db.Exec(sqlStatement, marketId, vendorId)
	if err != nil {
		fmt.Printf("Error: Could not remove vendor from market. %s", err)
		return "Error: Could not remove vendor from market.", err
	}
	return "Successfully removed vendor from market", nil
}

