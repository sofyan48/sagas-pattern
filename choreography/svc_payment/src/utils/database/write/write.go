package write

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// TransactionDB Global Connection ...
var TransactionDB *gorm.DB

// DBInit Initialization Connection
// return connection, error
func DBInit(dbhost, dbport, dbuser, dbname string) (*gorm.DB, error) {
	var (
		configDB = fmt.Sprintf(
			"postgresql://%s@%s:%s/%s?sslmode=disable",
			dbuser, dbhost, dbport, dbname)
	)

	DB, err := gorm.Open("postgres", configDB)
	if err != nil {
		log.Println(fmt.Sprintf("failed to connect to database: %v", err))
		return nil, err
	}
	return DB, nil
}

// GetTransactionConnection function
// return DB.Begin()
func GetTransactionConnection() *gorm.DB {
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	// dbpass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	if TransactionDB == nil {
		log.Println("No Active Connection Found")
		TransactionDB, _ = DBInit(dbhost, dbport, dbuser, dbname)
	}
	return TransactionDB
}
