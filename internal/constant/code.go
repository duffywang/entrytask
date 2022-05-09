package constant

//业务自定义状态码
var (
	//服务端异常
	ServerError  = NewError(100000, "Server Error")
	SessionError = NewError(100001, "Session Error")

	//客户端异常
	InvalidParamsError = NewError(200000, "Invalid Params Error")
	NotFoundError      = NewError(200001, "Not Found Error")

	//业务异常
	UserLoginError    = NewErrorWithData(300000, "User Login Error",[]string{"密码错误，请重试"})
	UserRegisterError = NewErrorWithData(300001, "User Register Error",[]string{"用户名已存在"})
	//UserRegisterError = NewErrorWithData(300001, "User Register Error",[]string{"用户名过长"})
	UserGetError      = NewErrorWithData(300002, "User Get Error",[]string{"SessionID错误"})
	UserEditError     = NewErrorWithData(300003, "User Edit Error",[]string{"SessionID错误"})
	FileUploadError   = NewError(300004, "File Upload Error")
	FileFormError     = NewError(300005, "File Form Error")
)
