package main

import (
	"log"
	"net/http"

	"github.com/canhlinh/ev-dict-server/app/http/routes"
	"github.com/canhlinh/ev-dict-server/app/services"
)

func main() {
	services.DbService()
	api := routes.MakeAPIRoutes()
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
