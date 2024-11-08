package repositories

import (
	"github.com/forzeyy/idea-shop-api/database"
	"github.com/forzeyy/idea-shop-api/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(models.Admin) (models.Admin, error)
	GetAdminByName(string) (models.Admin, error)
	GetAdminByID(uint) (models.Admin, error)
}

type adminRepository struct {
	conn *gorm.DB
}

func NewAdminRepository() AdminRepository {
	return &adminRepository{
		conn: database.ConnectDatabase(),
	}
}

func (db *adminRepository) CreateAdmin(admin models.Admin) (models.Admin, error) {
	return admin, db.conn.Create(&admin).Error
}

func (db *adminRepository) GetAdminByName(nickname string) (admin models.Admin, err error) {
	return admin, db.conn.Where("admin_name = ?", nickname).Find(&admin).Error
}

func (db *adminRepository) GetAdminByID(id uint) (admin models.Admin, err error) {
	return admin, db.conn.Where("id = ?", id).Find(&admin).Error
}
