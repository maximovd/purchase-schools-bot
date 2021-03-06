package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")

var dbinfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)

//CreateTable category table in database
func CreateTable() error {

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return err
	}
	defer db.Close()

	//Create category Table
	if _, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS purchase_categories(category_id SERIAL PRIMARY KEY, category_name TEXT NOT NULL);`,
	); err != nil {
		return err
	}

	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS purchases(
		purchase_id SERIAL PRIMARY KEY, 
		purchase_date DATE,
		purchase_name TEXT,
		purchase_count INT,
		category_id INT,
		CONSTRAINT fk_purchase
			FOREIGN KEY(category_id)
				REFERENCES purchase_categories(category_id)
		);`); err != nil {
		return err
	}

	return nil
}
