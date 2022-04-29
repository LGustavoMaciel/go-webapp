package routes

import (
	"net/http"
	"webApp/utils"
)

func homeGetHandler(w http.ResponseWriter, r *http.Request) {

	utils.ExecuteTemplate(w, "home.html", nil)

}