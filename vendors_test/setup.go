package vendors_test

import (
	"database/sql"
	"github.com/joho/godotenv"
	"os"
	"fmt"
	"strconv"
)

var Db *sql.DB

func SetupTestDB(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER")
	dbname := "market_money_test"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"dbname=%s sslmode=disable", host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("There was an error while connection to the database:", err)
		panic(err)
	} else {
		Db = db
		fmt.Println("Successfully connected to database")
	}
}

func TearDownTestDB() {
	Db.Query("DELETE * FROM vendors, markets, market_vendors")
}