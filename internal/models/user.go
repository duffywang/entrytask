package models

import (
	"gorm.io/gorm"
)

//用户
type User struct {
	*CommonModel
	UserName   string `json:"userName"`
	NickName   string `json:"nickName"`
	PassWord   string `json:"passWord"`
	ProfilePic string `json:"profile_pic"`
	Status     uint8  `json:"status"`
}



//工厂模式生成User
// func NewUser(username string, nickname string, picture string)User {
// 	return User{
// 		UserName : userName,
// 		NickName : nickName,
// 		ProfilePic : picture,
// 	}
// }

func (u User) CreateUserInfo(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (u User) UpdateUserInfo(db *gorm.DB, values any) error {
	err := db.Model(&u).Updates(values).Where("id = ? AND status = ?", u.ID, 0).Error
	if err != nil {
		return err
	}
	return nil
}

func (u User) GetUserInfoByName(db *gorm.DB) (*User, error) {
	var user User
	db = db.Where("name = ?", u.UserName)
	err := db.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
