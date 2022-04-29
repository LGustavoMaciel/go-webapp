package main

import (
	"fmt"
	"log"
	"net/http"
	"webApp/models"
	"webApp/routes"
	"webApp/utils"
)

const PORT = ":8080"

func main()  {
	models.TestConnection()
	fmt.Printf("Listening Port %s\n", PORT)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))

}

