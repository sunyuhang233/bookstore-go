package repository

import (
	"bookstore-go/global"
	"bookstore-go/model"
	"fmt"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO() *UserDAO {
	return &UserDAO{
		db: global.GetDB(),
	}
}

func (u *UserDAO) CreateUser(user *model.User) error {
	return u.db.Debug().Create(user).Error
}

func (u *UserDAO) CheckUserExists(username, phone, email string) (bool, error) {
	var count int64
	err := u.db.Model(&model.User{}).Where("username = ? OR phone = ? OR email = ?", username, phone, email).Count(&count).Error
	fmt.Println("CheckUserExists count:", count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (u *UserDAO) GetUserByUsername(username string) (*model.User, error) {
	var user *model.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDAO) GetUserByID(id int) (*model.User, error) {
	var user *model.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDAO) UpdateUser(user *model.User) error {
	return u.db.Save(user).Error
}
