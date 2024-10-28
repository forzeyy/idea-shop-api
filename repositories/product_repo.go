package repositories

import (
	"github.com/forzeyy/idea-shop-api/database"
	"github.com/forzeyy/idea-shop-api/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (models.Product, error)
	CreateProduct(models.Product) (models.Product, error)
	UpdateProduct(models.Product) (models.Product, error)
	DeleteProduct(models.Product) (models.Product, error)
}

type productRepository struct {
	conn *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		conn: database.ConnectDatabase(),
	}
}

func (db *productRepository) GetAllProducts() (products []models.Product, err error) {
	return products, db.conn.Find(&products).Error
}

func (db *productRepository) GetProductByID(id uint) (product models.Product, err error) {
	return product, db.conn.First(&product, id).Error
}

func (db *productRepository) CreateProduct(product models.Product) (models.Product, error) {
	return product, db.conn.Create(&product).Error
}

func (db *productRepository) UpdateProduct(product models.Product) (models.Product, error) {
	if err := db.conn.First(&product, product.ID).Error; err != nil {
		return product, err
	}
	return product, db.conn.Model(&product).Updates(&product).Error
}

func (db *productRepository) DeleteProduct(product models.Product) (models.Product, error) {
	if err := db.conn.First(&product, product.ID).Error; err != nil {
		return product, err
	}
	return product, db.conn.Delete(&product).Error
}
