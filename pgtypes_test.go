package pgtypes_test

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var sqlDB *sql.DB
var gormDB *gorm.DB

func init() {
	dsn := "host=localhost user=pgtypes password=pgtypes dbname=pgtypes sslmode=disable"

	if db, err := sql.Open("postgres", dsn); err != nil {
		panic(err)
	} else {
		sqlDB = db
	}

	if db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		gormDB = db
	}

	if err := gormDB.AutoMigrate(
		&tsvectorTestModel{},
	); err != nil {
		panic(err)
	}
}
