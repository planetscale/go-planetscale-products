package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Database *gorm.DB

func Connect() error {
	var err error
	dsn := fmt.Sprintf("%s&parseTime=True", os.Getenv("DSN"))

	Database, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{DisableForeignKeyConstraintWhenMigrating: true},
	)

	if err == nil {
		fmt.Println("Successfully connected to PlanetScale!")
	}

	return err
}
