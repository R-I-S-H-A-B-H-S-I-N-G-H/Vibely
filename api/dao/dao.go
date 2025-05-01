package dao

import (
	"errors"
	"fmt"

	databaseconfig "github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/database-config"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/dto"
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

// FindByID retrieves an entity by its ID
func (d *DAO[T]) FindByID(id string) (*T, error) {
	var entity T
	if err := d.db.First(&entity, "id", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil, nil for not found
		}
		return nil, err
	}
	return &entity, nil
}

func (d *DAO[T]) FindAll(filters map[string]interface{}, page, size int) (*dto.PaginationDTO[T], error) {
	var entities []T
	var totalCount int64

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * size
	query := d.db.Model(new(T))

	// Apply flexible filtering with partial matching
	for key, value := range filters {
		if value == nil {
			continue // Ignore missing fields
		}

		switch v := value.(type) {
		case string:
			query = query.Where(fmt.Sprintf("%s LIKE ?", key), "%"+v+"%") // Use LIKE for MySQL
		case int, int64, float64, bool:
			query = query.Where(fmt.Sprintf("%s = ?", key), v)
		case []interface{}:
			query = query.Where(fmt.Sprintf("%s IN (?)", key), v) // Support for lists
		default:
			query = query.Where(fmt.Sprintf("%s = ?", key), v)
		}
	}

	// Get total count
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	// Fetch paginated data
	if err := query.Limit(size).Offset(offset).Find(&entities).Error; err != nil {
		return nil, err
	}

	return &dto.PaginationDTO[T]{
		Data:       entities,
		TotalCount: totalCount,
		Page:       int64(page),
		Size:       int64(size),
	}, nil
}

func (d *DAO[T]) FindOne(query interface{}, args ...interface{}) (*T, error) {
	var entitiy T
	if err := d.db.Where(query, args...).First(&entitiy).Error; err != nil {
		return nil, err
	}
	return &entitiy, nil
}

// Update modifies an existing entity
func (d *DAO[T]) Update(entity *T) error {
	return d.db.Save(entity).Error
}

// Delete soft-deletes an entity by ID
func (d *DAO[T]) Delete(id string) error {
	var entity T
	if err := d.db.First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil // No error if already deleted/not found
		}
		return err
	}
	return d.db.Delete(&entity).Error
}

// DeletePermanent hard-deletes an entity by ID
func (d *DAO[T]) DeletePermanent(id string) error {
	var entity T
	if err := d.db.Unscoped().First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil // No error if not found
		}
		return err
	}
	return d.db.Unscoped().Delete(&entity).Error
}
func GetDAO[T any]() *DAO[T] {
	db := databaseconfig.GetMySqlDB()
	var entity T
	db.AutoMigrate(&entity)
	return &DAO[T]{db: db}
}
