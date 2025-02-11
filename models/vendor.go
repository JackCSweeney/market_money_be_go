package models

import (
	"example.com/mod/database"
	"fmt"
)

type Vendor struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ContactName    string `json:"contact_name"`
	ContactPhone   string `json:"contact_phone"`
	CreditAccepted bool   `json:"credit_accepted"`
}

func GetVendorById(id int) (Vendor, error) {
	sqlStatement := `SELECT * FROM vendors WHERE id=$1`
	var vendor Vendor

	row := database.Db.QueryRow(sqlStatement, id)
	err := row.Scan(&vendor.Id, &vendor.Name, &vendor.Description, &vendor.ContactName, &vendor.ContactPhone, &vendor.CreditAccepted)

	if err != nil {
		return Vendor{}, err
	}
	return vendor, nil
}

func GetAllVendors() ([]Vendor, error) {
	sqlStatement := `SELECT * FROM vendors`
	var vendors []Vendor

	rows, err := database.Db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var vendor Vendor
		err = rows.Scan(&vendor.Id, &vendor.Name, &vendor.Description, &vendor.ContactName, &vendor.ContactPhone, &vendor.CreditAccepted)
		if err != nil {
			fmt.Println("Error: Record missing needed information")
		}
		vendors = append(vendors, vendor)
	}
	err = rows.Err()
	if err != nil {
		return vendors, err
	}

	return vendors, nil
}