package api

import (
	"goji.io"
	"goji.io/pat"
)

var RootMux *goji.Mux

func InitRoute() {
	RootMux = goji.NewMux()
	RootMux.HandleFunc(pat.Get("/api/dictionary"), GetTranslateWord)
}
