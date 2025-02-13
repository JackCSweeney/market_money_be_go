package models

import (
	"example.com/mod/database"
	"fmt"
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
		fmt.Println(err)
		return Market{}, err
	}
	return market, nil
}

func GetAllMarkets() ([]Market, error) {
	sqlStatement := `SELECT * FROM markets`
	var markets []Market

	rows, err := database.Db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var market Market
		err = rows.Scan(&market.Id, &market.Name, &market.Street, &market.City, &market.County, &market.State, &market.Zip, &market.Lat, &market.Lon)
		if err != nil {
			fmt.Printf("Error: Record missing needed information: %s", err)
		}
		markets = append(markets, market)
	}
	err = rows.Err()
	if err != nil {
		return markets, err
	}
	return markets, nil
}