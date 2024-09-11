package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	template, err := template.ParseFiles(files...)
	if err != nil {
		app.errLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = template.Execute(w, nil)

}
