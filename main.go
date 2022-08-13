package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
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
	MainRouter.Handle("/images/{image}", http.StripPrefix("/images/", http.FileServer(http.Dir("assets/images"))))
	MainRouter.Handle("/fonts/{font}", http.StripPrefix("/fonts", http.FileServer(http.Dir("assets/fonts"))))
	MainRouter.Handle("/icon/{icon}", http.StripPrefix("/icon", http.FileServer(http.Dir("assets/icon"))))
	MainRouter.Handle("/{something}/css", http.StripPrefix("/css", http.FileServer(http.Dir("assets/css"))))
	// static files end
	MainRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles("src/pages/index.html", "src/parts/header.html", "src/parts/footer.html"))
		err := tmpl.ExecuteTemplate(writer, "index", nil)
		if err != nil {
			panic(err)
		}
	})
	MainRouter.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles("src/pages/index.html", "src/parts/header.html", "src/parts/footer.html"))
		err := tmpl.ExecuteTemplate(writer, "index", nil)
		if err != nil {
			panic(err)
		}
	})
	//Саб-роуты
	aboutSubRouter := MainRouter.PathPrefix("/about").Subrouter()
	aboutSubRouter.HandleFunc("", aboutIndex)
	aboutSubRouter.HandleFunc("/", aboutIndex)

	contactSubRouter := MainRouter.PathPrefix("/contact").Subrouter()
	contactSubRouter.HandleFunc("", contactIndex)
	contactSubRouter.HandleFunc("/", contactIndex)

	fashionSubRouter := MainRouter.PathPrefix("/fashion").Subrouter()
	fashionSubRouter.HandleFunc("", fashionIndex)
	fashionSubRouter.HandleFunc("/", fashionIndex)

	newsSubRouter := MainRouter.PathPrefix("/news").Subrouter()
	//newsSubRouter.HandleFunc("", newsIndex)
	newsSubRouter.HandleFunc("/", newsIndex)
	newsSubRouter.HandleFunc("/add", newsAddRoute)

	productsSubRouter := MainRouter.PathPrefix("/products").Subrouter()
	productsSubRouter.HandleFunc("", productsIndex)
	productsSubRouter.HandleFunc("/", productsIndex)
	productsSubRouter.HandleFunc("/add", productsAddRoute)

	uploadSubRouter := MainRouter.PathPrefix("/upload").Subrouter()
	uploadSubRouter.HandleFunc("", upload)
	uploadSubRouter.HandleFunc("/", upload)

	//Прослушивание портов
	err := http.ListenAndServe(":80", MainRouter)
	if err != nil {
		panic(err)
	}
}
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Метод:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
