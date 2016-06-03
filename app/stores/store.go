package stores

import (
	"time"

	"github.com/canhlinh/ev-dict-server/app/models"
	log "github.com/canhlinh/log4go"
)

type StoreResult struct {
	Data interface{}
	Err  *models.Error
}

type StoreChannel chan StoreResult

func Must(sc StoreChannel) interface{} {
	r := <-sc
	if r.Err != nil {
		log.Close()
		time.Sleep(time.Second)
		panic(r.Err)
	}
	return r.Data
}

type Store interface {
	Dict() DictStore
	Close()
}

type DictStore interface {
	Save(dict *models.DictEnVi) StoreChannel
	FindByWord(work string) StoreChannel
}
