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

func UpdateVendor(id int, updates Vendor) (Vendor, error){
	findStatement := `SELECT * FROM vendors WHERE id=$1`
	var vendorRow Vendor
	err := database.Db.QueryRow(findStatement, id).Scan(&vendorRow.Id, &vendorRow.Name, &vendorRow.Description, &vendorRow.ContactName, &vendorRow.ContactPhone, &vendorRow.CreditAccepted)
	if err != nil {
		return vendorRow, err
	}

	vendorRow = updateVendorName(vendorRow, updates.Name)
	vendorRow = updateVendorDescription(vendorRow, updates.Description)
	vendorRow = updateVendorContactName(vendorRow, updates.ContactName)
	vendorRow = updateVendorContactPhone(vendorRow, updates.ContactPhone)
	vendorRow = updateVendorCreditAccepted(vendorRow, updates.CreditAccepted)

	updateStatement := `UPDATE vendors SET name = $1, description = $2, contact_name = $3, contact_phone = $4, credit_accepted = $5 WHERE id=$6 RETURNING *;`
	var updatedVendor Vendor
	err = database.Db.QueryRow(updateStatement, vendorRow.Name, vendorRow.Description, vendorRow.ContactName, vendorRow.ContactPhone, vendorRow.CreditAccepted, id).Scan(&updatedVendor.Id, &updatedVendor.Name, &updatedVendor.Description, &updatedVendor.ContactName, &updatedVendor.ContactPhone, &updatedVendor.CreditAccepted)

	if err != nil {
		return updatedVendor, err
	}
	return updatedVendor, nil
}

func CreateVendor(name string, description string, contactName string, contactPhone string, creditAccepted bool) (Vendor, error){
	sqlStatement := `INSERT INTO vendors (name, description, contact_name, contact_phone, credit_accepted) VALUES ($1, $2, $3, $4, $5) RETURNING *;`
	newVendor := Vendor{}

	err := database.Db.QueryRow(sqlStatement, name, description, contactName, contactPhone, creditAccepted).Scan(&newVendor.Id, &newVendor.Name, &newVendor.Description, &newVendor.ContactName, &newVendor.ContactPhone, &newVendor.CreditAccepted)

	if err != nil {
		return newVendor, err
	}
	return newVendor, nil
}

func updateVendorName(vendor Vendor, name string) (Vendor) {
	if name != "" {
		vendor.Name = name
		return vendor
	}
	return vendor
}

func updateVendorDescription(vendor Vendor, description string) (Vendor) {
	if description != "" {
		vendor.Description = description
		return vendor
	}
	return vendor 
}

func updateVendorContactName(vendor Vendor, contactName string) (Vendor) {
	if contactName != "" {
		vendor.ContactName = contactName
		return vendor 
	}
	return vendor
}

func updateVendorContactPhone(vendor Vendor, contactPhone string) (Vendor) {
	if contactPhone != "" {
		vendor.ContactPhone = contactPhone
		return vendor 
	}
	return vendor
}

func updateVendorCreditAccepted(vendor Vendor, creditAccepted bool) (Vendor) {
	if creditAccepted != vendor.CreditAccepted {
		vendor.CreditAccepted = creditAccepted
		return vendor 
	}
	return vendor
}