package vendors

import (
	"fmt"
	"example.com/mod/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllVendors(c *gin.Context) {
	sqlStatement := `SELECT * FROM vendors`
	var data []Vendor

	rows, err := database.Db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var vendor Vendor
		err = rows.Scan(&vendor.ID, &vendor.Name, &vendor.Description, &vendor.ContactName, &vendor.ContactPhone, &vendor.CreditAccepted)
		if err != nil {
			fmt.Println("Error: Record missing needed information")
		}
		data = append(data, vendor)
	}
	err = rows.Err()
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Issue returning requested info")
	}

	c.JSON(http.StatusOK, data)
}