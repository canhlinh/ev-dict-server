package services

import (
	"log"
	"sync"

	"github.com/canhlinh/ev-dictionary/app/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type dbService struct {
	db *gorm.DB
}

var (
	dbServiceInstance *dbService
	once              sync.Once
)

func (this *dbService) InitDB() {
	var err error
	this.db, err = gorm.Open("mysql", "linhnc:123456@/vndict?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	this.db.LogMode(true)
	this.db.SingularTable(true)
	this.db.AutoMigrate(models.DictEnVi{})
	this.db.DB().SetMaxIdleConns(10)
	this.db.DB().SetMaxOpenConns(100)
}

func DbService() *gorm.DB {
	once.Do(func() {
		dbServiceInstance = &dbService{}
		dbServiceInstance.InitDB()
	})
	return dbServiceInstance.db
}
