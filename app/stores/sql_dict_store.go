package stores

import (
	"net/http"

	"github.com/canhlinh/ev-dict-server/app/models"
)

type SqlDictStore struct {
	*SqlStore
}

func NewSqlDictStore(sqlStore *SqlStore) DictStore {
	dictStore := &SqlDictStore{sqlStore}
	sqlStore.master.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.DictEnVi{})
	return dictStore
}

func (s *SqlDictStore) Save(dict *models.DictEnVi) StoreChannel {
	storeChannel := make(StoreChannel)
	return storeChannel
}

func (s *SqlDictStore) FindByWord(word string) StoreChannel {
	storeChannel := make(StoreChannel)
	go func() {
		result := StoreResult{}
		var dict models.DictEnVi
		if err := s.master.First(&dict, models.DictEnVi{Word: word}).Error; err != nil {
			result.Err = &models.Error{
				HttpCode:  http.StatusNotFound,
				ErrorCode: 1000,
				Message:   "Không tìm thấy dữ liệu tương ứng.",
			}
		} else {
			result.Data = dict
		}
		storeChannel <- result
	}()
	return storeChannel
}
