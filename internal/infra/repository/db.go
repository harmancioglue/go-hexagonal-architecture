package repository

import (
	"fmt"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/port/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type database struct {
	Database *gorm.DB
}

func NewDatabase() (repository.Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(db)

	return &database{
		Database: db,
	}, err
}

func (db *database) GetDatabase() *gorm.DB {
	return db.Database
}
