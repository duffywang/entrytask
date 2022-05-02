package dao

import (
	"fmt"
	"time"

	"github.com/duffywang/entrytask/internal/models"
)

//返回models.User 指针
func (d *Dao) CreateUser(userName, nickName, passWord, profilePic string, status uint8) (*models.User, error) {
	createUser := models.User{
		Username:   userName,
		Password:   passWord,
		Nickname:   nickName,
		ProfilePic: profilePic,
		Status:     status,

		CreateTime: uint32(time.Now().Unix()),
		UpdateTime: uint32(time.Now().Unix()),
	}

	user, err := createUser.CreateUserInfo(d.engine)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *Dao) UpdateUser(id uint32, nickName, profilePic string) error {
	//通过id查询到用户
	updateUser := models.User{
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

	err := updateUser.UpdateUserInfo(d.engine, values)
	return err
}

//TODO:没有返回指针
func (d *Dao) GetUserInfo(userName string) (models.User, error) {
	queryUser := models.User{Username: userName}
	user, err := queryUser.GetUserInfoByName(d.engine)
	if err != nil {
		fmt.Printf("Login.GetUserInfo.GetUserInfoByName Fail: %v \n", err)
		return models.User{}, err
	}
	return user, nil
}

//Update:删除用户信息
// func (d *Dao) DeleteUserInfo(userName string) (models.User, error) {

// }
