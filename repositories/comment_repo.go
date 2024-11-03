package repositories

import (
	"github.com/forzeyy/idea-shop-api/database"
	"github.com/forzeyy/idea-shop-api/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	GetAllComments() ([]models.Comment, error)
	GetCommentByID(id uint) (models.Comment, error)
	GetCommentsByUser(userID uint) ([]models.Comment, error)
	CreateComment(models.Comment) (models.Comment, error)
	UpdateComment(models.Comment) (models.Comment, error)
	DeleteComment(models.Comment) (models.Comment, error)
}

type commentRepository struct {
	conn *gorm.DB
}

func NewCommentRepository() CommentRepository {
	return &commentRepository{
		conn: database.ConnectDatabase(),
	}
}

func (db *commentRepository) GetAllComments() (comments []models.Comment, err error) {
	return comments, db.conn.Find(&comments).Error
}

func (db *commentRepository) GetCommentByID(id uint) (comment models.Comment, err error) {
	return comment, db.conn.First(&comment, id).Error
}

func (db *commentRepository) GetCommentsByUser(userID uint) (comments []models.Comment, err error) {
	return comments, db.conn.Where("user_id = ?", userID).Find(&comments).Error
}

func (db *commentRepository) CreateComment(comment models.Comment) (models.Comment, error) {
	return comment, db.conn.Create(&comment).Error
}

func (db *commentRepository) UpdateComment(comment models.Comment) (models.Comment, error) {
	var existingComment models.Comment
	if err := db.conn.First(&existingComment, comment.ID).Error; err != nil {
		return comment, err
	}
	return comment, db.conn.Model(&existingComment).Updates(comment).Error
}

func (db *commentRepository) DeleteComment(comment models.Comment) (models.Comment, error) {
	if err := db.conn.First(&comment, comment.ID).Error; err != nil {
		return comment, err
	}
	return comment, db.conn.Delete(&comment).Error
}
