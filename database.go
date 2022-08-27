package main

import (
	"database/sql"
	"fmt"
	"os"
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

var APP_LOGIN = os.Getenv("app_login")
var APP_PASSWORD = os.Getenv("app_password")
var APP_IP = os.Getenv("app_ip")
var APP_PORT = os.Getenv("app_port")
var APP_DBNAME = os.Getenv("app_dbname")

var db_path = fmt.Sprintf("%s:%s@(%s:%s)/%s", APP_LOGIN, APP_PASSWORD, APP_IP, APP_PORT, APP_DBNAME)

func ReturnDB() *sql.DB {

	db, err := sql.Open("mysql", db_path)
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

func AddingProduct(data NewProductDetails) {
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
	CreatedAt  time.Time
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
		err := rows.Scan(&one.id, &one.Title, &one.Text, &one.CreatedAt)
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
	CreatedAt   time.Time
	CreatedStr  string
}

func GetAllProdFromDB() []OneProdUnit {
	db := ReturnDB()
	// Все строки продуктов
	rows, err := db.Query(`SELECT id, description, price, image_path, created_at FROM products ORDER BY id DESC;`)
	defer rows.Close()
	check(err)
	var AllProdUnit []OneProdUnit
	for rows.Next() {
		var one OneProdUnit
		err := rows.Scan(&one.id, &one.Description, &one.Price, &one.ImagePath, &one.CreatedAt)
		check(err)
		AllProdUnit = append(AllProdUnit, one)
	}
	errAfter := rows.Err()
	check(errAfter)
	return AllProdUnit
}

// data for products
func CreateNewsData() []OneNewsUnit {
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
		err := rows.Scan(&one.id, &one.Title, &one.Text, &one.CreatedAt)
		check(err)
		AllNewsUnit = append(AllNewsUnit, one)
	}
	errAfter := rows.Err()
	check(errAfter)
	return AllNewsUnit
}

//data for products

// data for products
type DataForProd struct {
	ProdArray []OneProdUnit
}
type OneProdUnitFormat struct {
	id          int
	Description string
	Price       string
	ImagePath   string
	CreatedAt   time.Time
	CreatedStr  string
}
type DataForProdFormat struct {
	ProdArray []OneProdUnitFormat
}

func CreateProductsData() DataForProdFormat {

	dataFromDB := DataForProd{ProdArray: GetAllProdFromDB()}
	var MyData DataForProdFormat
	for _, unit := range dataFromDB.ProdArray {
		var row OneProdUnitFormat
		row.id = unit.id
		row.Description = unit.Description
		row.Price = unit.Price
		row.ImagePath = unit.ImagePath
		row.CreatedAt = unit.CreatedAt
		row.CreatedStr = unit.CreatedAt.Format("02-01-2006 3:04PM")
		MyData.ProdArray = append(MyData.ProdArray, row)
	}
	return MyData
}

//data for products

func CreateDBproducts() {
	db := ReturnDB()
	query := `
    CREATE TABLE products (
        id INT AUTO_INCREMENT,
        description TEXT NOT NULL,
        image_path TEXT NOT NULL,
        price FLOAT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
	_, err := db.Exec(query)
	check(err)
}
