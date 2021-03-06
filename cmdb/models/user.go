package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id           int64
	Name         string
	Password     string
	Birthday     *time.Time
	Telephone    string
	Email        string
	Addr         string
	Status       int8
	RoleId       int64
	DepartmentId int64
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
	Description  string
	Sex          bool
}

// 提供给 html 模板调用的方法，用于判断性别
func (u *User) SexText() string {
	if u.Sex {
		return "男"
	}
	return "女"
}

func (u *User) SetPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func (u *User) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
