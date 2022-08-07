package main

import (
	"html/template"
	"net/http"
	"time"
)

// About-page
func aboutIndex(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/pages/about.html"))
	err := tmpl.ExecuteTemplate(writer, "about", nil)
	if err != nil {
		panic(err)
	}
}

// Contact-page
func contactIndex(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/pages/contact.html"))
	err := tmpl.ExecuteTemplate(writer, "contact", nil)
	if err != nil {
		panic(err)
	}
}

// Fashion-page
func fashionIndex(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/pages/fashion.html"))
	err := tmpl.ExecuteTemplate(writer, "fashion", nil)
	if err != nil {
		panic(err)
	}
}

// News-page
func newsIndex(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/pages/news.html"))
	//type OneNewsUnit struct {
	//	id        int
	//	Title     string
	//	Text      string
	//	createdAt time.Time
	//}
	//var datass []OneUnitNews
	type DataForNews struct {
		NewsArray []OneNewsUnit
	}
	type OneNewsUnitFormat struct {
		id         int
		Title      string
		Text       string
		СreatedAt  time.Time
		СreatedStr string
	}
	type DataForNewsFormat struct {
		NewsArray []OneNewsUnitFormat
	}
	dataFromDB := DataForNews{NewsArray: GetAllNewsFromDB()}
	var Mydata DataForNewsFormat
	for _, unit := range dataFromDB.NewsArray {
		var row OneNewsUnitFormat
		row.id = unit.id
		row.Title = unit.Title
		row.Text = unit.Text
		row.СreatedAt = unit.СreatedAt
		row.СreatedStr = unit.СreatedAt.Format("02-01-2006 3:04PM")
		Mydata.NewsArray = append(Mydata.NewsArray, row)
	}
	//fmt.Println(Mydata)
	err := tmpl.ExecuteTemplate(writer, "news", Mydata)
	if err != nil {
		panic(err)
	}
}

// Products-page
func productsIndex(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/pages/products.html"))
	err := tmpl.ExecuteTemplate(writer, "products", nil)
	if err != nil {
		panic(err)
	}
}
