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
