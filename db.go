package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")

var dbinfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)

// PurchaseCategory struct for
type PurchaseCategory struct {
	ID   int
	name string
}

// CreateTables category and purchases tables in database
func CreateTables() error {

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return err
	}
	defer db.Close()

	// Create category Table
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

// GetAllCategory from database
func GetAllCategory() ([]int, []string, error) {
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return nil, nil, err
	}

	defer db.Close()

	var category PurchaseCategory
	var AllCategoriesName []string
	var AllCategoriesID []int

	categories, err := db.Query("SELECT * FROM purchase_categories;")
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	for categories.Next() {
		if err := categories.Scan(&category.ID, &category.name); err != nil {
			log.Fatal(err)
		}
		AllCategoriesID = append(AllCategoriesID, category.ID)
		AllCategoriesName = append(AllCategoriesName, category.name)
	}
	return AllCategoriesID, AllCategoriesName, err
}

//AddCategory into database
func AddCategory(category string) (string, error) {
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return "Ошибка подключения к базе", err
	}
	defer db.Close()
	query := "INSERT INTO purchase_categories(category_name) VALUES($1);"
	if _, err = db.Exec(query, category); err != nil {
		return "Ошибка выполнения запроса", err
	}
	return "Категория успешно добавлена", err
}

//DeleteCategory into database
func DeleteCategory(categoryID string) (string, error) {
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return "Ошибка подключения к базе", err
	}
	defer db.Close()
	query := "DELETE FROM purchase_categories WHERE category_id = $1;"
	if _, err = db.Exec(query, categoryID); err != nil {
		return "Ошибка выполнения запроса", err
	}
	return "Категория успешно удалена", err
}
