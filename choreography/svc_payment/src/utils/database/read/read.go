package read

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB Global Connetction
var DBRead *gorm.DB

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

// GetConnection function
// return connection
func GetConnection() *gorm.DB {
	dbhost := os.Getenv("DB_HOST_READ")
	dbport := os.Getenv("DB_PORT_READ")
	dbuser := os.Getenv("DB_USER_READ")
	// dbpass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME_READ")
	if DBRead == nil {
		log.Println("No Active Connection Found")
		DBRead, _ = DBInit(dbhost, dbport, dbuser, dbname)
	}
	return DBRead
}
