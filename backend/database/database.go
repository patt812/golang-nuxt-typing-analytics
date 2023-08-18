package database

import (
	"fmt"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/patt812/golang-nuxt-typing-analytics/domain"
	"github.com/patt812/golang-nuxt-typing-analytics/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func buildDSN() string {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	return fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbName)
}

func Connect() (*gorm.DB, error) {
	dsn := buildDSN()
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func Initialize() (*gorm.DB, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&domain.Kana{}, &domain.Pattern{})

	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations.GetAllMigrations())

	if err := m.Migrate(); err != nil {
		return nil, fmt.Errorf("could not migrate: %v", err)
	}

	return db, nil
}
