package models

import "github.com/jinzhu/gorm"

func Migrate(db *gorm.DB, appModels ...interface{}) *gorm.DB {
	db.AutoMigrate(appModels)

	return db
}
