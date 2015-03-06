package controllers

import (
    "net/http" 
    "html/template"
    "github.com/jasonmichels/out-of-office/page"  
)

type DashboardController struct {
}

func (controller *DashboardController) DashboardHandler(w http.ResponseWriter, r *http.Request) {

	var dashboardTemplate = template.Must(template.New("getDashboard").ParseFiles("templates/base.html", "templates/dashboard.html", "templates/topnav.html", "templates/footer.html"))
    p := &page.Page{Title: "Out Of Office Dashboard", ShowGoogleAnalytics: false, Cdn: "//localhost:8000", ServerUrl: "//localhost:9005"} 
    dashboardTemplate.ExecuteTemplate(w, "base", p)

}