package models

import (
	"gorm.io/gorm"

	"github.com/duffywang/entrytask/internal/constant"
)

//用户
type User struct {
	ID         uint32 `json:"id"`
	CreatedAt  uint32 `gorm:"column:create_time"`
	UpdatedAt  uint32 `gorm:"column:update_time"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Password   string `json:"password"`
	ProfilePic string `json:"profile_pic"`
	Status     uint8  `json:"status"`
}

/*
type User struct {
	gorm.Model
	Username string `gorm:"unique;type:varchar(50);not null;comment:用户名"`
	Nickname string `gorm:"type:varchar(50);comment:昵称"`
	Password string `gorm:"size:20;type:varchar(50);not null;comment:密码"`
	ProfilePic string `gorm:"column:profile_pic;comment:用户图片"`
	Status uint8 `gorm:"default:0;not null;comment:状态"`
}
*/

//创建用户
func (u User) CreateUserInfo(db *gorm.DB) (*User, error) {
	//Create(&u).ID Create(&u).RowsAffected
	err := db.Debug().Create(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

//更新用户信息
func (u User) UpdateUserInfo(db *gorm.DB, values any) (*User, error) {
	err := db.Debug().Model(&u).Updates(values).Where("id = ? AND status = ?", u.ID, 0).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

//使用设定的表名
func (u User) TableName() string {
	return constant.TableName
}

//通过用户名获取用户信息
func (u User) GetUserInfoByName(db *gorm.DB) (User, error) {
	var user User
	//临时使用Table("user_table")
	db = db.Where("username = ?", u.Username)
	err := db.Select("username", "password", "nickname", "profile_pic").Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

//删除用户
func (u User) DeleteUser(db *gorm.DB) error {
	var user User
	err := db.Where("username = ?", u.Username).Delete(&user).Error
	return err
}
