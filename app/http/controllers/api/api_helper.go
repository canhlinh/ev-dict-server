package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/canhlinh/ev-dict-server/app/models"
	"github.com/canhlinh/ev-dict-server/app/stores"
	"github.com/canhlinh/ev-dict-server/app/utils"
)

func SetupTest() {
	gopath := os.Getenv("GOPATH")
	configPathFile := gopath + "/src/github.com/canhlinh/ev-dict-server/conf/config.yaml"
	utils.LoadConfig(configPathFile)
	stores.NewMySQLStore()
	InitRoute()
}

func RenderJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func RenderErrorJSON(w http.ResponseWriter, err *models.Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HttpCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(err)
}
