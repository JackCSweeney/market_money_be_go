package marketvendors

import (
	"encoding/json"
	"fmt"
	"example.com/mod/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMarketVendor(c *gin.Context) {
	body := MarketVendor{}
	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(400, "MarketVendor is not defined")
		return
	}

	err = json.Unmarshal(data, &body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	sqlStatement := `INSERT INTO market_vendors(vendor_id,market_id) RETURNING id`
	id := 0
	
	err = database.Db.QueryRow(sqlStatement, body.VendorID, body.MarketID).Scan(&id)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, "Could not create new Market Vendor")
	}

	c.JSON(http.StatusCreated, "Successfully added vendor to market")
}