package database

import (
	"fmt"

	"github.com/wellmtx/challenge/internal/infra/entities"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB         *gorm.DB
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	Test       bool
}

func NewDatabase(dbHost, dbUser, dbPassword, dbName string, test bool) *Database {
	return &Database{
		DBHost:     dbHost,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		Test:       test,
	}
}

func (d *Database) Connect() error {
	fmt.Println("Connecting to database...")

	if !d.Test {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", d.DBHost, d.DBUser, d.DBPassword, d.DBName)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("failed to connect to database: %w", err)
		}
		d.DB = db
	} else {
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("failed to connect to database: %w", err)
		}
		d.DB = db

		if err := d.DB.AutoMigrate(&entities.GyroscopeEntity{}, &entities.GeolocationEntity{}, &entities.PhotoEntity{}); err != nil {
			return fmt.Errorf("failed to migrate test database: %w", err)
		}
	}
	fmt.Println("Connected to database")

	return nil
}

func (d *Database) Ping() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	fmt.Println("Closed database connection")

	return nil
}
