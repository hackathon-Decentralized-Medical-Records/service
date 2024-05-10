package mysql

import (
	"gorm.io/gorm"
	"service/utils/httputils"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"type:int;not null" json:"id" `
	UserName string `gorm:"type:varchar(100);not null" json:"username" `
	PassWord string `gorm:"type:varchar(65);not null" json:"password" `
	Email    string `gorm:"type:varchar(100);not null" json:"email" `
	Role     int    `gorm:"type:int" json:"role"`
}

func (User) TableName() string {
	return "user"
}

// 校验用户是否存在
func CheckUserByEmail(email string) int {
	var user User
	db.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		return httputils.StatusConflict
	}
	return httputils.StatusOK

}

// 校验登录信息
func GetUserByNameAndPwd(loginId, passWord string) (int, string) {
	var user User
	db.Where("email = ? AND pass_word = ?", loginId, passWord).First(&user)
	if user.ID <= 0 {
		return httputils.StatusNotFound, "用户不存在"
	}
	return httputils.StatusOK, "登录成功"
}

// 注册用户信息
func RegisterUser(user *User) (int, string) {
	db.Create(user)
	return httputils.StatusOK, "注册成功"
}