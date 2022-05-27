package model

import (
	"fmt"
	"gorm.io/gorm"
	"readCommunity/global"
)


type User struct {
	gorm.Model
	UserName    string `gorm:"column:user_name; " json:"user_name"`
	NickName    string `gorm:"column:nick_name;" json:"nick_name"`
	Password    string `gorm:"column:password" json:"password"`
	Description string `gorm:"column:description" json:"description"`
	Email       string `gorm:"column:email" json:"email"`
	Phone       string `gorm:"column:phone" json:"phone"`
	Avatar      string `gorm:"column:avatar" json:"avatar"`
	RoleId      int    `gorm:"column:role_id" json:"role_id"`
	Status      uint8  `gorm:"column:status" json:"status"`
	CreateId    int    `gorm:"column:create_id" json:"create_id"`
	//LastLoginTime time.Time `gorm:"column:last_login_time" json:"last_login_time"`
}

func (u User) TableName() string {
	return "rd_user"
}

// 注册用户
func (u User) AddUser(user User) error {
	return global.DB.Create(&user).Error
}

func (u User)CheckAuth(username, password string) (bool, error) {
	var user User
	fmt.Println("model.user:>>>", username, password)
	err := global.DB.Where(User{UserName: username, Password: password}).First(&user).Error
	if err != nil {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}
// 验证用户名、密码正确，登陆
func (u User) CheckPwd(username, password string) (isTrue bool, err error) {
	var count int64
	err = global.DB.Model(u).Where("user_name=? and password = ?", username, password).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		isTrue = true
	} else {
		isTrue = false
	}
	return
}
