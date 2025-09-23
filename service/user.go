package service

import (
	"bookstore-go/model"
	"bookstore-go/repository"
	"encoding/base64"
	"errors"
)
type UserService struct {
	UserDB *repository.UserDAO
}

func NewUserService() *UserService {
	return &UserService{
		UserDB: repository.NewUserDAO(),
	}
}


func (u *UserService) Register(username, password, email, phone, confirmPassword string) error {
	exists, err := u.UserDB.CheckUserExists(username, phone, email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("用户已存在")
	}
	if password != confirmPassword {
		return errors.New("两次输入的密码不一致")
	}
	encodePassword:=u.encodePassword(password)
	err = u.createUser(username, encodePassword, phone, email)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) encodePassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}


func (u *UserService) createUser(username,password,phone,email string) error {
	user:=model.User{
		Username: username,
		Password: password,
		Phone: phone,
		Email: email,
	}
	return u.UserDB.CreateUser(&user)
}
