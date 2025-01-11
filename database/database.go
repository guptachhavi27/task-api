package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	postgresdb "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	fmt.Printf("dsn is: %s\n", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance from GORM:", err)
	}

	runMigrations( sqlDB)

	DB = db
}

func runMigrations( sqlDB *sql.DB) {
	driver, err := postgresdb.WithInstance(sqlDB, &postgresdb.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}
	
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", 
		"postgres",         
		driver,
	)
	if err != nil {
		log.Fatal("Failed to initialize migrations:", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	} else {
		log.Println("Migrations applied successfully.")
	}
}
