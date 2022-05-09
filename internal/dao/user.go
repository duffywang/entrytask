package dao

import (
	"fmt"
	"time"

	"github.com/duffywang/entrytask/internal/constant"
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
func (d *Dao) UpdateUser(id uint32, nickName, profilePic string) (*models.User, error) {
	//通过id查询到用户
	u := models.User{
		ID:        id,
		UpdatedAt: uint32(time.Now().Unix()),
	}

	values := map[string]any{}
	if nickName != "" {
		values[constant.Nickname] = nickName
	}

	if profilePic != "" {
		values[constant.ProfilePic] = profilePic
	}

	user, err := u.UpdateUserInfo(d.engine, values)
	return user,err
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
func (d *Dao) DeleteUserInfo(userName string) error {
	u := models.User{Username: userName}
	err := u.DeleteUser(d.engine)
	return err
}
