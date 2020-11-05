package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username       string `json:"username"`
	PasswordDigest string `json:"password_digest"`
	Nickname       string `json:"nickname"`
	Avatar         string `json:"avatar"`
	// 个性签名
	Signature string `json:"signature"`
}

const (
	// 密码加密难度 4 ~ 31
	PasswordCost = 15
)

func GetUserById(id interface{}) (User, error) {
	var user User
	err := DB.Where("id = ?", id).First(&user).Error
	return user, err
}

// 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
