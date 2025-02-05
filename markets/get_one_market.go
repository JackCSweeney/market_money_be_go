package markets

import (
	"fmt"
	"strconv"
	"database/sql"
	"example.com/mod/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOneMarket(c *gin.Context, id int) {
	sqlStatement := `SELECT * FROM markets WHERE id=$1`
	var market Market

	row := database.Db.QueryRow(sqlStatement, id)
	err := row.Scan(&market.ID, &market.Name, &market.Street, &market.City, &market.County, &market.State, &market.Zip, &market.Lat, &market.Lon)

	errorMessage := fmt.Sprintf("Error: Market record not found with id: %s", strconv.Itoa(id))

	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusNotFound, errorMessage)
	case nil:
		c.JSON(http.StatusOK, market)
	default:
		c.JSON(http.StatusBadRequest, err)
	}
}