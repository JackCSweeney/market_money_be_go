package vendors

import (
	"database/sql"
	"example.com/mod/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOneVendor(c *gin.Context, id int) {
	sqlStatment := `SELECT * FROM vendors WHERE id=$1`
	var vendor Vendor

	row := database.Db.QueryRow(sqlStatment, id)
	err := row.Scan(&vendor.ID, &vendor.Name, &vendor.Description, &vendor.ContactName, &vendor.ContactPhone, &vendor.CreditAccepted)

	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusNotFound, "Error: Record not found")
	case nil:
		c.JSON(http.StatusOK, vendor)
	default:
		c.JSON(http.StatusBadRequest, err)
	}
}