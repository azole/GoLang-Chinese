package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path --> 抓第一個斜線後的值，包括斜線
	// http://localhost:4000/azole/test/ha
	// r.URL.Path = /azole/test/ha
	// r.URL.Path[1:] --> azole/test/ha

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	p, err := loadPage(title)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	// if p, err := loadPage(title); err == nil{
	//   fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	// }
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	// 用 FormValue 拿到的都是 string
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/view/{title}", viewHandler)
	r.HandleFunc("/edit/{title}", editHandler)
	r.HandleFunc("/save/{title}", saveHandler)
	http.Handle("/", r)

	http.ListenAndServe(":4000", nil)
}
