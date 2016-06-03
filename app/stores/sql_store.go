package stores

import (
	"github.com/canhlinh/ev-dict-server/app/utils"
	log "github.com/canhlinh/log4go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var MysqlStore *SqlStore

type SqlStore struct {
	master *gorm.DB
	dict   DictStore
}

func initConnection() *SqlStore {
	sqlStore := &SqlStore{}
	dbConnection, err := gorm.Open("mysql", utils.GetConfig().SqlConfig.DataSource)
	if err != nil {
		log.Crashf("Got error when connect database, the error is '%v'", err)
	}
	dbConnection.SingularTable(true)
	dbConnection.DB().SetMaxIdleConns(10)
	dbConnection.DB().SetMaxOpenConns(100)
	dbConnection.LogMode(utils.GetConfig().SqlConfig.Debug)
	dbConnection.SetLogger(log.GormLogger{})
	sqlStore.master = dbConnection
	return sqlStore
}

func NewMySQLStore() {
	MysqlStore = initConnection()
	MysqlStore.dict = NewSqlDictStore(MysqlStore)
}

func (s SqlStore) Close() {
	s.master.Close()
}

func (s SqlStore) Dict() DictStore {
	return s.dict
}
