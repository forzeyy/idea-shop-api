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
	GetProductsByCategory(categoryID uint) ([]models.Product, error)
	UpdateProductImageURL(id uint, imageURL string) error
	SearchProducts(query string) ([]models.Product, error)
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
	var existingProduct models.Product
	if err := db.conn.First(&existingProduct, product.ID).Error; err != nil {
		return product, err
	}
	return product, db.conn.Model(&existingProduct).Updates(product).Error
}

func (db *productRepository) DeleteProduct(product models.Product) (models.Product, error) {
	if err := db.conn.First(&product, product.ID).Error; err != nil {
		return product, err
	}
	return product, db.conn.Delete(&product).Error
}

func (db *productRepository) GetProductsByCategory(categoryID uint) (products []models.Product, err error) {
	return products, db.conn.Joins("JOIN product_categories ON product_categories.id = products.id").Where("product_categories.category_id = ?", categoryID).Find(&products).Error
}

func (db *productRepository) UpdateProductImageURL(id uint, imageURL string) error {
	return db.conn.Model(&models.Product{}).Where("id = ?", id).Update("image_url", imageURL).Error
}

func (db *productRepository) SearchProducts(query string) (products []models.Product, err error) {
	return products, db.conn.Where("name LIKE ? OR description LIKE ?", "%"+query+"%").Find(&products).Error
}
