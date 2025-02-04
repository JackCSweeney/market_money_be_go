package vendors

import (
	"database/sql"
	"example.com/mod/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOneVendor(c *gin.Context, id int) {
	sqlStatement := `SELECT * FROM vendors WHERE id=$1`
	var vendor Vendor

	row := database.Db.QueryRow(sqlStatement, id)
	err := row.Scan(&vendor.ID, &vendor.Name, &vendor.Description, &vendor.ContactName, &vendor.ContactPhone, &vendor.CreditAccepted)

	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusNotFound, "Error: Vendor record not found")
	case nil:
		c.JSON(http.StatusOK, vendor)
	default:
		c.JSON(http.StatusBadRequest, err)
	}
}