package vendors

import (
	"encoding/json"
	"example.com/mod/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateVendor(c *gin.Context) {
	body := Vendor{}
	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	err = json.Unmarshal(data, &body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
	}

	sqlStatement := `UPDATE vendors SET name = $1, description = $2, contact_name = $3, contact_phone = $4, credit_accepted = $5 WHERE id=$6 RETURNING *`
	var vendor Vendor

	errUpdate := database.Db.QueryRow(sqlStatement, body.Name, body.Description, body.ContactName, body.ContactPhone, body.CreditAccepted, body.ID).Scan(&vendor.ID, &vendor.Name, &vendor.Description, &vendor.ContactName, &vendor.ContactPhone, &vendor.CreditAccepted)
	if errUpdate != nil {
		c.JSON(400, "Error: Issue updating record for vendor")
		panic(errUpdate)
	}

	c.JSON(http.StatusOK, vendor)
}