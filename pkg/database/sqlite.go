package database

import (
	"sync"

	"github.com/h4zlq/cli/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	once     sync.Once
	instance *gorm.DB
)

func GetInstance() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		db.AutoMigrate(&model.Task{})
		instance = db
	})

	return instance
}
