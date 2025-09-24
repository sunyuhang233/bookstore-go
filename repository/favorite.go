package repository

import (
	"bookstore-go/global"
	"bookstore-go/model"

	"gorm.io/gorm"
)

type FavoriteDAO struct {
	db *gorm.DB
}

func NewFavoriteDAO() *FavoriteDAO {
	return &FavoriteDAO{
		db: global.GetDB(),
	}
}

func (dao *FavoriteDAO) AddFavorite(userId, bookId int) error {
	return dao.db.Create(&model.Favorite{
		UserID: userId,
		BookID: bookId,
	}).Error
}

func (dao *FavoriteDAO) RemoveFavorite(userId, bookId int) error {
	return dao.db.Where("user_id = ? AND book_id = ?", userId, bookId).Delete(&model.Favorite{}).Error
}

func (dao *FavoriteDAO) GetFavorites(userId, page, pageSize int) ([]*model.Favorite, int64, error) {
	var favorites []*model.Favorite
	var total int64
	err := dao.db.Preload("Book").Where("user_id = ?", userId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&favorites).Error
	if err != nil {
		return nil, 0, err
	}
	err = dao.db.Model(&model.Favorite{}).Where("user_id = ?", userId).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return favorites, total, err
}
