package utilities

import (
	"fmt"
	"log"
	"strconv"

	"../models"
	gorm "github.com/jinzhu/gorm"

	// Get the guts for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// Hold on to the database connection for all queries
var db *gorm.DB

// Hold on to the error
var err error

// Init creates a connection to a postgres database and returns the DB instance
func Init(cfg *Configuration) {
	user := getEnvironmentVariable("PG_USER", cfg.DB.User)
	password := getEnvironmentVariable("PG_PASSWORD", cfg.DB.Password)
	host := getEnvironmentVariable("PG_HOST", cfg.DB.Host)
	port := getEnvironmentVariable("PG_PORT", strconv.Itoa(cfg.DB.Port))
	database := getEnvironmentVariable("PG_DB", cfg.DB.Database)

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	Connect(dbinfo)
}

// Connect will connect to the database and return a DB instance
func Connect(dbinfo string) {
	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")
}

// BuildDatabase sets up the database if there are no tables
func BuildDatabase(db *gorm.DB) {
	if !db.HasTable(&models.SampleModel{}) {
		err := db.CreateTable(&models.SampleModel{})
		if err != nil {
			log.Println("Table already exists")
		}
	}
	db.AutoMigrate(&models.SampleModel{})
}

// Close closes the database connection
func Close(db *gorm.DB) {
	db.Close()
}
