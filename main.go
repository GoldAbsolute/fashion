package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func main() {
	//Основной роут
	MainRouter := mux.NewRouter()
	// static files start
	// classical http static
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// mux path route
	MainRouter.Handle("/css/{style}", http.StripPrefix("/css", http.FileServer(http.Dir("assets/css"))))
	MainRouter.Handle("/js/{script}", http.StripPrefix("/js", http.FileServer(http.Dir("assets/js"))))
	MainRouter.Handle("/images/{image}", http.StripPrefix("/images/", http.FileServer(http.Dir("assets/images/"))))
	MainRouter.Handle("/fonts/{font}", http.StripPrefix("/fonts", http.FileServer(http.Dir("assets/fonts"))))
	MainRouter.Handle("/icon/{icon}", http.StripPrefix("/icon", http.FileServer(http.Dir("assets/icon"))))
	// static files end
	MainRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles("src/pages/index.html", "src/parts/header.html", "src/parts/footer.html"))
		err := tmpl.ExecuteTemplate(writer, "index", nil)
		if err != nil {
			panic(err)
		}
	})
	//Саб-роуты
	aboutSubRouter := MainRouter.PathPrefix("/about").Subrouter()
	aboutSubRouter.HandleFunc("/", aboutIndex)

	contactSubRouter := MainRouter.PathPrefix("/contact").Subrouter()
	contactSubRouter.HandleFunc("/", contactIndex)

	fashionSubRouter := MainRouter.PathPrefix("/fashion").Subrouter()
	fashionSubRouter.HandleFunc("/", fashionIndex)

	newsSubRouter := MainRouter.PathPrefix("/news").Subrouter()
	newsSubRouter.HandleFunc("/", newsIndex)

	productsSubRouter := MainRouter.PathPrefix("/products").Subrouter()
	productsSubRouter.HandleFunc("/", productsIndex)

	//Прослушивание портов
	err := http.ListenAndServe(":80", MainRouter)
	if err != nil {
		panic(err)
	}
}
