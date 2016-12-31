package web

import (
	"fmt"
	"html/template"
	"net/http"
)

type DashboardController struct{}

func NewDashboardController() *DashboardController {
	return &DashboardController{}
}

type Context struct {
	Title string
}

func (dc *DashboardController) Index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	t, err := t.ParseFiles("web/template/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Execute(w, Context{Title: "Dashboard here"})
}
