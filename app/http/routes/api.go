package routes

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
	api_controllers "github.com/canhlinh/ev-dict-server/app/http/controllers/api"
)

func MakeAPIRoutes() *rest.Api {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/api/dictionary", api_controllers.GetTranslateWord),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	return api
}
