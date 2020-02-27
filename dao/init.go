package dao

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var orm *gorm.DB

func init() {
	var onceGorm sync.Once
	onceGorm.Do(func() {
		var err error
		orm, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/aaron?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}
		if !orm.HasTable(&TodoModel{}) {
			orm.CreateTable(&TodoModel{})
		}
		//迁移the schema
		orm.AutoMigrate(&TodoModel{})
	})
}
