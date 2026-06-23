package adventure_time

import (
	"html/template"
	"net/http"
)

type Options struct {
	Text string
	arc  string
}

type PageData struct {
	PageTitle string
	PageBody  string
	//PageOptions Options
}

func StoryHandler(writer http.ResponseWriter, req *http.Request) {
	data := PageData{
		PageTitle: "My First Go Template",
		PageBody:  "Welcome to My Website",
		//Message: "This basic page was rendered dynamically using Go's html/template package!",
	}
	template, err := template.ParseFiles("./adventure_time/storytime.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	template.Execute(writer, data)
}
