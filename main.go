package main

import (
	"time"

	"github.com/canhlinh/ev-dict-server/app/api"
	"github.com/canhlinh/ev-dict-server/app/stores"
	"github.com/canhlinh/ev-dict-server/app/utils"
	log "github.com/canhlinh/log4go"
	"github.com/tylerb/graceful"
)

func main() {
	utils.LoadConfig("./conf/config.yaml")
	stores.NewMySQLStore()
	api.InitRoute()
	api.RootMux.Use(log.NewGojiLog())
	log.Info("Start EV dictionary on listen port : " + utils.GetConfig().ListenAddress)
	graceful.Run(utils.GetConfig().ListenAddress, 5*time.Second, api.RootMux)
}
