package repositories

import (
	"github.com/forzeyy/idea-shop-api/database"
	"github.com/forzeyy/idea-shop-api/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders() ([]models.Order, error)
	GetOrdersByUser(userID uint) ([]models.Order, error)
	GetOrderByID(id uint) (models.Order, error)
	CreateOrder(models.Order) (models.Order, error)
	UpdateOrder(models.Order) (models.Order, error)
	DeleteOrder(models.Order) (models.Order, error)
	AcceptOrder(models.Order) (models.Order, error)
}

type orderRepository struct {
	conn *gorm.DB
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{
		conn: database.ConnectDatabase(),
	}
}

func (db *orderRepository) GetAllOrders() (orders []models.Order, err error) {
	return orders, db.conn.Find(&orders).Error
}

func (db *orderRepository) GetOrdersByUser(userID uint) (orders []models.Order, err error) {
	return orders, db.conn.Where("user_id = ?", userID).Find(&orders).Error
}

func (db *orderRepository) GetOrderByID(id uint) (order models.Order, err error) {
	return order, db.conn.First(&order, id).Error
}

func (db *orderRepository) CreateOrder(order models.Order) (models.Order, error) {
	return order, db.conn.Create(&order).Error
}

func (db *orderRepository) UpdateOrder(order models.Order) (models.Order, error) {
	var existingOrder models.Order
	if err := db.conn.First(&existingOrder, order.ID).Error; err != nil {
		return order, err
	}
	return order, db.conn.Model(&existingOrder).Updates(order).Error
}

func (db *orderRepository) DeleteOrder(order models.Order) (models.Order, error) {
	if err := db.conn.First(&order, order.ID).Error; err != nil {
		return order, err
	}
	return order, db.conn.Delete(&order).Error
}

func (db *orderRepository) AcceptOrder(order models.Order) (models.Order, error) {
	return order, db.conn.Raw("UPDATE orders SET is_accepted = ? WHERE id = ?", true, order.ID).Error
}

func (db *orderRepository) MakeOrderInactive(order models.Order) (models.Order, error) {
	return order, db.conn.Raw("UPDATE orders SET is_active = ? WHERE id = ?", false, order.ID).Error
}
