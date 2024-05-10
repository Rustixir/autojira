package infrastructure

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"smart/domain/model"
)

func createTables(db *gorm.DB) error {
	return db.AutoMigrate(
		new(model.Issue),
	)
}

func NewDBClient() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("autojira"), new(gorm.Config))
	if err != nil {
		return nil, err
	}

	if err = createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}
