package vendors

import (
	"encoding/json"
	"fmt"
	"database/sql"
	"example.com/mod/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Vendor struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ContactName    string `json:"contact_name"`
	ContactPhone   string `json:"contact_phone"`
	CreditAccepted bool   `json:"credit_accepted"`
}

func AddVendor(c *gin.Context) {
	body := Vendor{}
	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(400, "Vendor is not defined")
		return
	}

	err = json.Unmarshal(data, &body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	sqlStatement := `INSERT INTO vendors(name,description,contact_name,contact_phone,credit_accepted) values ($1,$2,$3,$4,$5) RETURNING id`
	id := 0

	err = database.Db.QueryRow(sqlStatement, body.Name, body.Description, body.ContactName, body.ContactPhone, body.CreditAccepted).Scan(&id)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, "Could not create new vendor")
	}

	sqlStatementTwo := `SELECT * FROM vendors WHERE id=$1`
	var vendor Vendor

	row := database.Db.QueryRow(sqlStatementTwo, id)
	errFind := row.Scan(&vendor.ID, &vendor.Name, &vendor.Description, &vendor.ContactName, &vendor.ContactPhone, &vendor.CreditAccepted)
	switch errFind {
	case sql.ErrNoRows:
		fmt.Println("No vendor was returned")
		return
	case nil:
		c.JSON(http.StatusCreated, vendor)
	default:
		panic(errFind)
	}
}