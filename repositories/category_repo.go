package repositories

import (
	"github.com/forzeyy/idea-shop-api/database"
	"github.com/forzeyy/idea-shop-api/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (models.Category, error)
	CreateCategory(models.Category) (models.Category, error)
	UpdateCategory(models.Category) (models.Category, error)
	DeleteCategory(models.Category) (models.Category, error)
}

type categoryRepository struct {
	conn *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		conn: database.ConnectDatabase(),
	}
}

func (db *categoryRepository) GetAllCategories() (categories []models.Category, err error) {
	return categories, db.conn.Find(&categories).Error
}

func (db *categoryRepository) GetCategoryByID(id uint) (category models.Category, err error) {
	return category, db.conn.First(&category, id).Error
}

func (db *categoryRepository) CreateCategory(category models.Category) (models.Category, error) {
	return category, db.conn.Create(&category).Error
}

func (db *categoryRepository) UpdateCategory(category models.Category) (models.Category, error) {
	var existingCategory models.Category
	if err := db.conn.First(&existingCategory, category.ID).Error; err != nil {
		return category, err
	}
	return category, db.conn.Model(&existingCategory).Updates(category).Error
}

func (db *categoryRepository) DeleteCategory(category models.Category) (models.Category, error) {
	if err := db.conn.First(&category, category.ID).Error; err != nil {
		return category, err
	}
	return category, db.conn.Delete(&category).Error
}
