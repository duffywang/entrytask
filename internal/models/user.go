package model
import "fmt"

//用户
type User struct {
	*CommonModel
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	PassWord string `json:"passWord"`
	ProfilePic string `json:"profile_pic"`
	Status uint8 `json:"status"`
}

//工厂模式生成User
func NewUser(username string, nickname string, picture string)User {
	return User{
		UserName : userName, 
		NickName : nickName, 
		ProfilePic : picture,
	}
}

func (u User)Create(db *gorm.DB) (*User, err){
	db.Create()
}



func (this User)GetUserInfo() string{
	userInfo := fmt.Sprintf("%v\t%v\t%v\t",this.Username,this.Nickname,this.Picture)
	return userInfo
}