package dao

//返回models.User 指针
func (d *Dao)CreateUser(userName,nickName,passWord,profilePic string,status int)(*models.User,error){

}


func (d *Dao)UpdateUser(userName,nickName,profilePic string)(error){

}

//TODO:没有返回指针
func (d *Dao)GetUserInfo(userName string) (models.User, error) {
	
}

//Update:删除用户信息
func (d *Dao)DeleteUserInfo(userName string)(models.User, error){

}