package db

import (
	logger "LibraryManagementSystem/log"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var database *sql.DB

func Connect() {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSource := fmt.Sprintf(dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName)
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err.Error())
	}

	database = db

	// Try to ping the database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to the database")
}

func GetDatabase() *sql.DB {
	return database
}

func Migrate() {

	/*dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSource := fmt.Sprintf("sql://%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("DB Source is:", dbSource)
	migrationFileSource := "file://./migration/migration.sql"

	m, err := migrate.New(migrationFileSource, dbSource)
	if err != nil {
		panic(err.Error())
	}*/
	logger.Log("Migration completed")
}
