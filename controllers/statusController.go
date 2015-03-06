package controllers

import (
    "net/http" 
    "html/template"
    "github.com/jasonmichels/out-of-office/page"  
)

type StatusController struct {
}

func (controller *StatusController) CreateHandler(w http.ResponseWriter, r *http.Request) {

	var createStatusTemplate = template.Must(template.New("createStatus").ParseFiles("templates/base.html", "templates/status/create.html", "templates/topnav.html", "templates/footer.html"))
    p := &page.Page{Title: "Out Of Office Create Status", ShowGoogleAnalytics: false, Cdn: "//localhost:8000", ServerUrl: "//localhost:9005"} 
    createStatusTemplate.ExecuteTemplate(w, "base", p)

}