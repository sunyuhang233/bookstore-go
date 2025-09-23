package service

import (
	"bookstore-go/jwt"
	"bookstore-go/model"
	"bookstore-go/repository"
	"encoding/base64"
	"errors"
	"fmt"
)

type UserService struct {
	UserDB *repository.UserDAO
}

type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int64     `json:"expires_in"`
	UserInfo     *UserInfo `json:"user_info"`
}

type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
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
	encodePassword := u.encodePassword(password)
	err = u.createUser(username, encodePassword, phone, email)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Login(username, password, captcha string, captchaID string) (*LoginResponse, error) {
	encodePassword := u.encodePassword(password)
	user, err := u.UserDB.GetUserByUsername(username)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	if user.Password != encodePassword {

		return nil, errors.New("密码错误")
	}

	token, err := jwt.GenerateTokenPair(uint(user.ID), user.Username)
	if err != nil {
		return nil, err
	}
	userInfo := &UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Avatar:   user.Avatar,
	}
	return &LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
		UserInfo:     userInfo,
	}, nil
}

func (u *UserService) GetUserInfo(userId int) (*UserInfo, error) {
	user, err := u.UserDB.GetUserByID(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	return &UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Avatar:   user.Avatar,
	}, nil
}

func (u *UserService) UpdateUserInfo(userId int, username, email, phone, avatar string) (*model.User, error) {
	user, err := u.UserDB.GetUserByID(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	user.Username = username
	user.Email = email
	user.Phone = phone
	user.Avatar = avatar
	update_err := u.UserDB.UpdateUser(user)
	if update_err != nil {
		return nil, update_err
	}
	return u.UserDB.GetUserByID(userId)
}

func (u *UserService) ChangePassword(userId int, oldPassword, newPassword string) error {
	user, err := u.UserDB.GetUserByID(userId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("用户不存在")
	}
	encodeOldPassword := u.encodePassword(oldPassword)
	if user.Password != encodeOldPassword {
		return errors.New("旧密码错误")
	}
	if user.Password == u.encodePassword(newPassword) {
		return errors.New("新密码不能与旧密码相同")
	}
	encodeNewPassword := u.encodePassword(newPassword)
	user.Password = encodeNewPassword
	return u.UserDB.UpdateUser(user)
}

func (u *UserService) encodePassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}

func (u *UserService) createUser(username, password, phone, email string) error {
	user := model.User{
		Username: username,
		Password: password,
		Phone:    phone,
		Email:    email,
	}
	return u.UserDB.CreateUser(&user)
}
