package main

import (
	"database/sql"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func ConnectDatabase() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/http_db?parseTime=true")
	check(err)
	err2 := db.Ping()
	check(err2)
}
func ReturnDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/http_db?parseTime=true")
	check(err)
	err2 := db.Ping()
	check(err2)
	return db
}
func CreateDB() {
	//writer http.ResponseWriter, request *http.Request
	db := ReturnDB()
	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
	_, err := db.Exec(query)
	check(err)

	//_, err3 := fmt.Fprint(writer, "База данных успешно создана!")
	//check(err3)
}
func CreateTableNews() {
	//writer http.ResponseWriter, request *http.Request
	db := ReturnDB()
	query := `
    CREATE TABLE news (
        id INT AUTO_INCREMENT,
        title TEXT NOT NULL,
        text TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
	_, err := db.Exec(query)
	check(err)

	//_, err3 := fmt.Fprint(writer, "База данных успешно создана!")
	//check(err3)
}
func CreateDefaultNews() {
	db := ReturnDB()
	title := "default_title_text"
	text := "_not_default_text_text_for_news_in_row_"
	createdAt := time.Now()
	_, err := db.Exec(`INSERT INTO news (title, text, created_at) VALUES (?, ?, ?)`, title, text, createdAt)
	check(err)
	//liid64, err := result.LastInsertId()
	//check(err)
	//ra64, err2 := result.RowsAffected()
	//check(err2)
	//_, err3 := fmt.Fprintf(writer, "Последнее добавленое Id(сейчас добавилось): %d \nЧисло затронутых строк: %d \nПользователь %s успешно создан!\n", liid64, ra64, username)
	//check(err3)
}
func AddingNews(data NewsDetails) {
	db := ReturnDB()
	name := data.Name
	email := data.Email
	_, _ = name, email
	title := data.Title
	text := data.Text
	//createdAt := data.CreatedAt
	createdAt := time.Now()
	_, err := db.Exec(`INSERT INTO news (title, text, created_at) VALUES (?, ?, ?)`, title, text, createdAt)
	check(err)
}

func AddingProduct(data NewsProductDetails) {
	db := ReturnDB()
	//name := data.Name
	//email := data.Email
	//_, _ = name, email
	description := data.Description
	price := data.Price
	imagePath := data.ImagePath
	createdAt := time.Now()
	_, err := db.Exec(`INSERT INTO products (description, price, image_path, created_at) VALUES (?, ?, ?, ?)`, description, price, imagePath, createdAt)
	check(err)
}

type OneNewsUnit struct {
	id         int
	Title      string
	Text       string
	СreatedAt  time.Time
	CreatedStr string
}

func GetAllNewsFromDB() []OneNewsUnit {
	db := ReturnDB()
	// Все строки новостей

	rows, err := db.Query(`SELECT id, title, text, created_at FROM news ORDER BY id DESC;`)
	defer rows.Close()
	check(err)
	var AllNewsUnit []OneNewsUnit
	//type ArrayOfNewsStruct struct {
	//	ArrayOfNews []OneNewsUnit
	//}
	//var AllNewsUnit = ArrayOfNewsStruct{}
	for rows.Next() {
		var one OneNewsUnit
		err := rows.Scan(&one.id, &one.Title, &one.Text, &one.СreatedAt)
		check(err)
		AllNewsUnit = append(AllNewsUnit, one)
	}
	errAfter := rows.Err()
	check(errAfter)
	return AllNewsUnit
}

type OneProdUnit struct {
	id          int
	Description string
	Price       string
	ImagePath   string
	СreatedAt   time.Time
	CreatedStr  string
}

func GetAllProdFromDB() []OneProdUnit {
	db := ReturnDB()
	// Все строки продуктов
	rows, err := db.Query(`SELECT id, description, price, image_path, created_at FROM news ORDER BY id DESC;`)
	defer rows.Close()
	check(err)
	var AllProdUnit []OneProdUnit
	for rows.Next() {
		var one OneProdUnit
		err := rows.Scan(&one.id, &one.Description, &one.Price, &one.ImagePath, &one.СreatedAt)
		check(err)
		AllProdUnit = append(AllProdUnit, one)
	}
	errAfter := rows.Err()
	check(errAfter)
	return AllProdUnit
}
func CreateDBproducts() {
	db := ReturnDB()
	query := `
    CREATE TABLE products (
        id INT AUTO_INCREMENT,
        description TEXT NOT NULL,
        price FLOAT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
	_, err := db.Exec(query)
	check(err)
}
