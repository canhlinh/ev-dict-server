package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/canhlinh/ev-dict-server/app/models"
	"github.com/canhlinh/ev-dict-server/app/stores"
)

func GetTranslateWord(w rest.ResponseWriter, r *rest.Request) {
	word := r.FormValue("word")
	if word == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(models.Error{
			ErrorCode: 1000,
			Message:   "Thiếu parameter",
		})
		return
	}

	result := <-stores.MysqlStore.Dict().FindByWord(word)
	if result.Err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.WriteJson(models.Error{
			ErrorCode: 1001,
			Message:   "Không tìm thấy dữ liệu",
		})
		return
	}
	dict := result.Data.(models.DictEnVi)
	w.WriteJson(dict)
}
