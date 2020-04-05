package database

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB Global Connetction
var DB *gorm.DB

// TransactionDB Global Connection ...
var TransactionDB *gorm.DB

// DBInit Initialization Connection
// return connection, error
func DBInit() (*gorm.DB, error) {
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	// dbpass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	var (
		configDB = flag.String("addr", fmt.Sprintf(
			"postgresql://%s@%s:%s/%s?sslmode=disable",
			dbuser, dbhost, dbport, dbname), "DB SETUP")
	)

	DB, err := gorm.Open("postgres", *configDB)
	if err != nil {
		log.Println(fmt.Sprintf("failed to connect to database: %v", err))
		return nil, err
	}
	return DB, nil
}

// GetConnection function
// return connection
func GetConnection() *gorm.DB {
	if DB == nil {
		log.Println("No Active Connection Found")
		DB, _ = DBInit()
	}
	return DB
}

// GetTransactionConnection function
// return DB.Begin()
func GetTransactionConnection() *gorm.DB {
	if TransactionDB == nil {
		log.Println("No Active Connection Found")
		TransactionDB, _ = DBInit()
	}
	return TransactionDB
}
