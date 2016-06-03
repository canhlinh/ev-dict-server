package main

import (
	"log"
	"net/http"

	"github.com/canhlinh/ev-dict-server/app/http/routes"
	"github.com/canhlinh/ev-dict-server/app/stores"
	"github.com/canhlinh/ev-dict-server/app/utils"
)

func main() {
	utils.LoadConfig("./conf/config.yaml")
	stores.NewMySQLStore()
	api := routes.MakeAPIRoutes()
	log.Fatal(http.ListenAndServe(":8000", api.MakeHandler()))
}
