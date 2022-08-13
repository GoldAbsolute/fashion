package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// About-page
func aboutIndex(writer http.ResponseWriter, request *http.Request) {
	_ = request
	tmpl := template.Must(template.ParseFiles("src/pages/about.html"))
	err := tmpl.ExecuteTemplate(writer, "about", nil)
	if err != nil {
		panic(err)
	}
}

// Contact-page
func contactIndex(writer http.ResponseWriter, request *http.Request) {
	_ = request
	tmpl := template.Must(template.ParseFiles("src/pages/contact.html"))
	err := tmpl.ExecuteTemplate(writer, "contact", nil)
	if err != nil {
		panic(err)
	}
}

// Fashion-page
func fashionIndex(writer http.ResponseWriter, request *http.Request) {
	_ = request
	tmpl := template.Must(template.ParseFiles("src/pages/fashion.html"))
	err := tmpl.ExecuteTemplate(writer, "fashion", nil)
	if err != nil {
		panic(err)
	}
}

// News-page
func newsIndex(writer http.ResponseWriter, request *http.Request) {
	_ = request
	tmpl := template.Must(template.ParseFiles("src/pages/news.html"))
	type DataForNews struct {
		NewsArray []OneNewsUnit
	}
	type OneNewsUnitFormat struct {
		id         int
		Title      string
		Text       string
		CreatedAt  time.Time
		CreatedStr string
	}
	type DataForNewsFormat struct {
		NewsArray []OneNewsUnitFormat
	}
	dataFromDB := DataForNews{NewsArray: GetAllNewsFromDB()}
	var MyData DataForNewsFormat
	for _, unit := range dataFromDB.NewsArray {
		var row OneNewsUnitFormat
		row.id = unit.id
		row.Title = unit.Title
		row.Text = unit.Text
		row.CreatedAt = unit.CreatedAt
		row.CreatedStr = unit.CreatedAt.Format("02-01-2006 3:04PM")
		MyData.NewsArray = append(MyData.NewsArray, row)
	}
	//fmt.Println(MyData)
	err := tmpl.ExecuteTemplate(writer, "news", MyData)
	if err != nil {
		panic(err)
	}
}

type NewsDetails struct {
	Name      string
	Email     string
	Title     string
	Text      string
	CreatedAt time.Time
}

func newsAddRoute(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		details := NewsDetails{
			Name:      request.FormValue("Name"),
			Email:     request.FormValue("Email"),
			Title:     request.FormValue("Title"),
			Text:      request.FormValue("Text"),
			CreatedAt: time.Now(),
		}
		AddingNews(details)
		http.Redirect(writer, request, "/news/", http.StatusSeeOther)
	}
	tmpl := template.Must(template.ParseFiles("src/pages/news_add.html"))
	err := tmpl.ExecuteTemplate(writer, "news_add", nil)
	if err != nil {
		panic(err)
	}
}

// Products-page
func productsIndex(writer http.ResponseWriter, request *http.Request) {
	_ = request
	tmpl := template.Must(template.ParseFiles("src/pages/products.html"))

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
	dataFromDB := DataForProd{ProdArray: GetAllProdFromDB()}
	var Mydata DataForProdFormat
	for _, unit := range dataFromDB.ProdArray {
		var row OneProdUnitFormat
		row.id = unit.id
		row.Description = unit.Description
		row.Price = unit.Price
		row.ImagePath = unit.ImagePath
		row.CreatedAt = unit.CreatedAt
		row.CreatedStr = unit.CreatedAt.Format("02-01-2006 3:04PM")
		Mydata.ProdArray = append(Mydata.ProdArray, row)
	}

	err := tmpl.ExecuteTemplate(writer, "products", Mydata)
	if err != nil {
		panic(err)
	}
}

type NewProductDetails struct {
	Name        string
	Email       string
	Description string
	Price       string
	PriceFloat  float64
	ImagePath   string
	CreatedAt   time.Time
}

func productsAddRoute(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		err01 := request.ParseMultipartForm(32 << 20)
		check(err01)
		file, handler, err := request.FormFile("Image")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_ = handler
		filePath := "./assets/images/products/" + handler.Filename
		f, errF := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if errF != nil {
			panic(errF)
		}
		defer f.Close()
		_, err04 := io.Copy(f, file)
		check(err04)
		//fmt.Println(filePath)
		filePath = strings.Replace(filePath, "./", "/", 1)
		//fmt.Println(filePath)
		details := NewProductDetails{
			Name:        request.FormValue("Name"),
			Email:       request.FormValue("Email"),
			Description: request.FormValue("Description"),
			Price:       request.FormValue("Price"),
			PriceFloat:  1,
			ImagePath:   filePath,
			CreatedAt:   time.Now(),
		}
		_ = details
		fmt.Println(details)
		AddingProduct(details)
		http.Redirect(writer, request, "/products/", http.StatusSeeOther)
	}
	tmpl := template.Must(template.ParseFiles("src/pages/products_add.html"))
	err := tmpl.ExecuteTemplate(writer, "products_add", nil)
	if err != nil {
		panic(err)
	}
}
