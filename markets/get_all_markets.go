package markets

import (
	"fmt"
	"example.com/mod/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllMarkets(c *gin.Context) {
	sqlStatement := `SELECT * FROM markets`
	var data []Market 

	rows, err := database.Db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var market Market
		err = rows.Scan(&market.ID, &market.Name, &market.Street, &market.City, &market.County, &market.State, &market.Zip, &market.Lat, &market.Lon)

		if err != nil {
			fmt.Println("Error: Record missing needed information")
		}

		data = append(data, market)
	}
	err = rows.Err()
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Issue returning requested information")
	}

	c.JSON(http.StatusOK, data)
}