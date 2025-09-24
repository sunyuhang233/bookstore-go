package repository

import (
	"bookstore-go/global"
	"bookstore-go/model"

	"gorm.io/gorm"
)

type BookDAO struct {
	db *gorm.DB
}

func NewBookDAO() *BookDAO {
	return &BookDAO{
		db: global.GetDB(),
	}
}

func (b *BookDAO) GetHotBooks(limit int) ([]model.Book, error) {
	var books []model.Book
	err := b.db.Where("status = ?", 1).Order("sale DESC").Limit(limit).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *BookDAO) GetNewBooks(limit int) ([]model.Book, error) {
	var books []model.Book
	err := b.db.Where("status = ?", 1).Order("created_at DESC").Limit(limit).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *BookDAO) GetBooks(page, pageSize int) ([]*model.Book, int, error) {
	var books []*model.Book
	var total int64
	offset := (page - 1) * pageSize

	err := b.db.Where("status = ?", 1).Offset(offset).Limit(pageSize).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}
	b.db.Model(&model.Book{}).Where("status = ?", 1).Count(&total)
	return books, int(total), nil
}

func (b *BookDAO) SearchBooks(page int, pageSize int, query string) ([]*model.Book, int, error) {
	var books []*model.Book
	var total int64
	offset := (page - 1) * pageSize

	searchCondition := b.db.Where("status = ? AND (title LIKE ? OR author LIKE ? OR description LIKE ?)", 1, "%"+query+"%", "%"+query+"%", "%"+query+"%")
	err := searchCondition.Model(&model.Book{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = searchCondition.Offset(offset).Limit(pageSize).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}
	return books, int(total), nil
}

func (b *BookDAO) GetBookByID(id uint) (*model.Book, error) {
	var book model.Book
	err := b.db.Where("id = ? AND status = ?", id, 1).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}
