package repositories

import (
	"github.com/forzeyy/idea-shop-api/database"
	"github.com/forzeyy/idea-shop-api/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	GetUserByPhone(phone string) (models.User, error)
	CreateUser(models.User) (models.User, error)
	UpdateUser(models.User) (models.User, error)
	DeleteUser(models.User) (models.User, error)
}

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		conn: database.ConnectDatabase(),
	}
}

func (db *userRepository) GetAllUsers() (users []models.User, err error) {
	return users, db.conn.Find(&users).Error
}

func (db *userRepository) GetUserByID(id uint) (user models.User, err error) {
	return user, db.conn.First(&user, id).Error
}

func (db *userRepository) GetUserByPhone(phone string) (user models.User, err error) {
	return user, db.conn.Where("phone = ?", phone).First(&user).Error
}

func (db *userRepository) CreateUser(user models.User) (models.User, error) {
	return user, db.conn.Create(&user).Error
}

func (db *userRepository) UpdateUser(user models.User) (models.User, error) {
	var existingUser models.User
	if err := db.conn.First(&existingUser, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.conn.Model(&existingUser).Updates(user).Error
}

func (db *userRepository) DeleteUser(user models.User) (models.User, error) {
	if err := db.conn.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.conn.Delete(&user).Error
}
