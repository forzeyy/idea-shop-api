package repositories

import (
	"github.com/forzeyy/idea-shop-api/database"
	"github.com/forzeyy/idea-shop-api/models"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetCartByUserID(userID uint) (models.Cart, error)
	AddItemToCart(cartID uint, item models.CartItem) (models.CartItem, error)
	UpdateCartItem(cartID uint, item models.CartItem) (models.CartItem, error)
	RemoveCartItem(cartID uint, item models.CartItem) (models.CartItem, error)
	ClearCart(cartID uint) error
}

type cartRepository struct {
	conn *gorm.DB
}

func NewCartRepository() CartRepository {
	return &cartRepository{
		conn: database.ConnectDatabase(),
	}
}

func (db *cartRepository) GetCartByUserID(userID uint) (cart models.Cart, err error) {
	return cart, db.conn.Preload("Items").Where("user_id = ?", userID).First(&cart).Error
}

func (db *cartRepository) AddItemToCart(cartID uint, item models.CartItem) (models.CartItem, error) {
	item.CartID = cartID
	return item, db.conn.Create(&item).Error
}

func (db *cartRepository) UpdateCartItem(cartID uint, item models.CartItem) (models.CartItem, error) {
	item.CartID = cartID
	var existingItem models.CartItem
	if err := db.conn.First(&existingItem, item.ID).Error; err != nil {
		return item, err
	}

	return item, db.conn.Model(&existingItem).Updates(item).Error
}

func (db *cartRepository) RemoveCartItem(cartID uint, item models.CartItem) (models.CartItem, error) {
	return item, db.conn.Where("cart_id = ? AND id = ?", cartID, item.ID).Delete(&item).Error
}

func (db *cartRepository) ClearCart(cartID uint) error {
	return db.conn.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error
}
