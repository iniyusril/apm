package db

import (
	"apm/model"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
)

var dbInternal *gorm.DB

func Init() *gorm.DB {

	dsn := `host=xxx user=xxx password=xxx dbname=xxx port=xxx sslmode=xxx TimeZone=xxx`
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Blog{})

	dbInternal = db
	return dbInternal
}

func GetDb() *gorm.DB {
	return dbInternal
}
