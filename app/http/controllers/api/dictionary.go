package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/canhlinh/ev-dictionary/app/http/responses"
	"github.com/canhlinh/ev-dictionary/app/models"
	"github.com/canhlinh/ev-dictionary/app/services"
)

func GetTranslateWord(w rest.ResponseWriter, r *rest.Request) {
	word := r.FormValue("word")
	if word == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(responses.Error{
			ErrorCode: 1000,
			Message:   "Thiếu parameter",
		})
		return
	}
	var result models.DictEnVi
	if err := services.DbService().First(&result, models.DictEnVi{Word: word}).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.WriteJson(responses.Error{
			ErrorCode: 1001,
			Message:   "Không tìm thấy dữ liệu",
		})
		return
	}
	w.WriteJson(result)
}
