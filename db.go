package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	type User struct {
		gorm.Model
		id       int    `json:"id" gorm:"primary_key"`
		name     string `json:"name"`
		email    string `json:"email"`
		password string `json:"password"`
	}

	type Role struct {
		gorm.Model
		id   int    `json:"id" gorm:"primary_key"`
		name string `json:"name"`
	}

	type Group struct {
		gorm.Model
		id      int    `json:"id" gorm:"primary_key"`
		name    string `json:"name"`
		members string `json:"members"`
	}

	db.AutoMigrate(&User{}, &Role{}, &Group{})

}
