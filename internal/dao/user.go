package dao

import (
	"github.com/duffywang/entrytask/internal/models"
)

//返回models.User 指针
func (d *Dao) CreateUser(userName, nickName, passWord, profilePic string, status int) (*models.User, error) {
	queryUser := models.User{UserName: userName}
	user,err :=  queryUser.GetUserInfoByNamed(d.engine)
	if err != nil {
		return nil,err
	}
	return user,nil
}

func (d *Dao) UpdateUser(userName, nickName, profilePic string) error {

}

//TODO:没有返回指针
func (d *Dao) GetUserInfo(userName string) (models.User, error) {

}

//Update:删除用户信息
func (d *Dao) DeleteUserInfo(userName string) (models.User, error) {

}
