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