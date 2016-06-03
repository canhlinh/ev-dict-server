package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/canhlinh/ev-dict-server/app/models"
	"github.com/canhlinh/gojit"
	"github.com/stretchr/testify/assert"
)

func TestGetTranslateWord(t *testing.T) {
	SetupTest()
	r := gojit.New()
	r.GET("/api/dictionary?word=love").
		Run(RootMux, func(r gojit.HTTPResponse, rq gojit.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			data := []byte(r.Body.String())
			var dict models.DictEnVi
			err := json.Unmarshal(data, &dict)
			assert.Equal(t, nil, err)
			assert.Equal(t, "love", dict.Word)
			assert.Equal(t, "lʌv", dict.Phonetic)
		})
}

func TestGetTranslatePluralWord(t *testing.T) {
	SetupTest()
	r := gojit.New()
	r.GET("/api/dictionary?word=potatoes").
		Run(RootMux, func(r gojit.HTTPResponse, rq gojit.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			data := []byte(r.Body.String())
			var dict models.DictEnVi
			err := json.Unmarshal(data, &dict)
			assert.Equal(t, nil, err)
			assert.Equal(t, "potato", dict.Word)
			assert.Equal(t, "pə'teitou", dict.Phonetic)
		})
}
