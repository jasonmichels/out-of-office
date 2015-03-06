package controllers

import (
    "net/http" 
    "html/template"
    "github.com/jasonmichels/out-of-office/page"
    "github.com/jasonmichels/out-of-office/auth"    
)

type AccountController struct {
}

func (controller *AccountController) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var loginTemplate = template.Must(template.New("getLogin").ParseFiles("templates/base.html", "templates/login.html", "templates/topnav.html", "templates/footer.html"))
    p := &page.Page{Title: "Out Of Office", ShowGoogleAnalytics: false, Cdn: "//localhost:8000", ServerUrl: "//localhost:9005"} 
    loginTemplate.ExecuteTemplate(w, "base", p)

}

func (controller *AccountController) AuthenticateHandler(w http.ResponseWriter, r *http.Request) {

	auth := &auth.Auth{}
	var isAuthenticated = auth.Attempt(r.FormValue("username"), r.FormValue("password"))

	if isAuthenticated {
		// set a flash message and redirect to their profile page or intended page
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		// set the flash message for error and return to the login page
		// also possible set an error in the session
		http.Redirect(w, r, "/login", http.StatusFound)
	}

}