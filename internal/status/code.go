package status

//业务自定义状态码
var (
	//服务端异常
	ServerError  = NewError(10000, "Server Error")
	SessionError = NewError(10001, "Session Error")

	//客户端异常
	InvalidParamsError = NewError(20000, "Invalid Params Error")
	NotFoundError      = NewError(20001, "Not Found Error")

	//业务异常
	UserLoginError    = NewError(30000, "User Login Error")
	UserRegisterError = NewError(30001, "User Register Error")
	UserGetError      = NewError(30002, "User Get Error")
	UserEditError     = NewError(30003, "User Edit Error")
	FileUploadError   = NewError(30004, "File Upload Error")
)
