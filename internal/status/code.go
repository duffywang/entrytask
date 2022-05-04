package status

//业务自定义状态码
var (

	//服务端异常
	ServerError  = NewError(100000, "Server Error")
	SessionError = NewError(100001, "Session Error")

	//客户端异常
	InvalidParamsError = NewError(200000, "Invalid Params Error")
	NotFoundError      = NewError(200001, "Not Found Error")

	//业务异常
	UserLoginError    = NewError(300000, "User Login Error")
	UserRegisterError = NewError(300001, "User Register Error")
	UserGetError      = NewError(300002, "User Get Error")
	UserEditError     = NewError(300003, "User Edit Error")
	FileUploadError   = NewError(300004, "File Upload Error")
	FileFormError     = NewError(300005, "File Form Error")
)