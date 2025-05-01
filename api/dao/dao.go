package dao

import (
	databaseconfig "github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/database-config"
	"gorm.io/gorm"
)

type DAO[T any] struct {
	db *gorm.DB
}

// Create inserts a new entity into the database and returns the saved entity
func (d *DAO[T]) Create(entity *T) (*T, error) {
	err := d.db.Create(entity).Error
	if err != nil {
		return nil, err // Return nil entity and the error if save fails
	}
	return entity, nil // Return the saved entity and nil error on success
}

func GetDAO[T any]() *DAO[T] {
	db := databaseconfig.GetMySqlDB()
	var entity T
	db.AutoMigrate(&entity)
	return &DAO[T]{db: db}
}
