package models

import (
	"fmt"

	"gorm.io/gorm"
)

//用户
type User struct {
	ID         uint32 `json:"id"`
	CreateTime uint32 `json:"create_time,omitempty" `
	UpdateTime uint32 `json:"update_time,omitempty"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Password   string `json:"password"`
	ProfilePic string `json:"profile_pic"`
	Status     uint8  `json:"status"`
}

/*
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Nickname string `gorm:"type:varchar(50)"`
	Password string `gorm:"size:20;type:varchar(50);not null"`
	ProfilePic string `gorm:"column:profile_pic"`
	Status uint8 `gorm:"default:0;not null"`
	Ignore string `gorm:"-"`
}
*/



func (u User) CreateUserInfo(db *gorm.DB) (*User, error) {
	//Create(&u).ID Create(&u).RowsAffected
	err := db.Table("user_table").Create(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (u User) UpdateUserInfo(db *gorm.DB, values any) error {
	err := db.Table("user_table").Model(&u).Updates(values).Where("id = ? AND status = ?", u.ID, 0).Error
	if err != nil {
		return err
	}
	return nil
}

//实现TableName接口
func (u User) TableName() string {
	return "user_table"
}

func (u User) GetUserInfoByName(db *gorm.DB) (User, error) {
	var user User
	//临时使用"user_table"
	db = db.Table("user_table").Where("username = ?", u.Username)
	err := db.Table("user_table").First(&user).Error
	//db = db.Table("user_table").Find(&user, "username = ?", u.Username)
	fmt.Printf("Login.GetUserInfo.GetUserInfoByName User is : %+v\n", user)
	if err != nil {
		return user, err
	}
	return user, nil
}
