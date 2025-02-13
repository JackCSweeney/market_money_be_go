package models

import (
	"example.com/mod/database"
	// "fmt"
)

type Market struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Street string `json:"street"`
	City string `json:"city"`
	County string `json:"county"`
	State string `json:"state"`
	Zip string `json:"zip"`
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func GetMarketById(id int) (Market, error) {
	sqlStatement := `SELECT * FROM markets WHERE id=$1`
	var market Market

	err := database.Db.QueryRow(sqlStatement, id).Scan(&market.Id, &market.Name, &market.Street, &market.City, &market.County, &market.State, &market.Zip, &market.Lat, &market.Lon)

	if err != nil {
		return Market{}, err
	}
	return market, nil
}

// func GetAllMarkets() ([]Market, error) {

// }