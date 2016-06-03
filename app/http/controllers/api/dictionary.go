package api

import (
	"net/http"
	"strings"

	"github.com/canhlinh/ev-dict-server/app/models"
	"github.com/canhlinh/ev-dict-server/app/stores"
	"github.com/jinzhu/inflection"
	"golang.org/x/net/context"
)

func GetTranslateWord(c context.Context, w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")
	if word == "" {
		RenderErrorJSON(w, &models.Error{
			HttpCode:  http.StatusBadRequest,
			ErrorCode: 1000,
			Message:   "Thiếu parameter",
		})
		return
	}

	word = strings.ToLower(word)
	result := <-stores.MysqlStore.Dict().FindByWord(word)
	if result.Err != nil {
		singularWord := inflection.Singular(word)
		if singularWord == word {
			RenderErrorJSON(w, result.Err)
			return
		}
		result = <-stores.MysqlStore.Dict().FindByWord(singularWord)
		if result.Err != nil {
			RenderErrorJSON(w, result.Err)
			return
		}
	}
	dict := result.Data.(models.DictEnVi)
	RenderJSON(w, dict)
}
