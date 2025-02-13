package database

import(
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

var Db *sql.DB

func ConnectDatabase(nameKey string) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error has occurred with loading .env file")
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER")
	dbname := os.Getenv(nameKey)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"dbname=%s sslmode=disable", host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("There was an error while connecting to the database", err)
		panic(err)
	} else {
		Db = db
		fmt.Println("Successfully connected to database")
	}
}