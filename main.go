package main

import (
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	gorm.Model
	Title string
	Price int
}

type Result struct {
	Total int
}

// DB migration
func dbInit() {
	db, err := gorm.Open("sqlite3", "book.sqlite3")
	if err != nil {
		panic("You can't open DB (dbInit())")
	}
	defer db.Close()
	db.AutoMigrate(&Book{})
