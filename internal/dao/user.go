package dao

import (
	"fmt"
	"time"

	"github.com/duffywang/entrytask/internal/models"
)

//DAO 创建用户
func (d *Dao) CreateUser(userName, nickName, passWord, profilePic string, status uint8) (*models.User, error) {
	u := models.User{
		Username:   userName,
		Password:   passWord,
		Nickname:   nickName,
		ProfilePic: profilePic,
		Status:     status,

		CreatedAt: uint32(time.Now().Unix()),
		UpdatedAt: uint32(time.Now().Unix()),
	}

	user, err := u.CreateUserInfo(d.engine)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//DAO 更新用户信息
func (d *Dao) UpdateUser(id uint32, nickName, profilePic string) error {
	//通过id查询到用户
	u := models.User{
		ID: id,
	}

	values := map[string]any{
		"UpdateTime": uint32(time.Now().Unix()),
	}
	if nickName != "" {
		values["nickname"] = nickName
	}

	if profilePic != "" {
		values["profile_pic"] = profilePic
	}

	err := u.UpdateUserInfo(d.engine, values)
	return err
}

//DAO 获取用户信息
func (d *Dao) GetUserInfo(userName string) (models.User, error) {
	u := models.User{Username: userName}
	user, err := u.GetUserInfoByName(d.engine)
	if err != nil {
		fmt.Printf("Login.GetUserInfo.GetUserInfoByName Fail: %v \n", err)
		return models.User{}, err
	}
	return user, nil
}

//DAO 删除用户
 func (d *Dao) DeleteUserInfo(userName string)  error {
	u := models.User{Username: userName}
	err := u.DeleteUser(d.engine)
	return err
 }
