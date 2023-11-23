package code

import "MyRainFarm/pkg/xrrcode"

var (
	RegisterMobileEmpty   = xrrcode.New(100001, "注册手机号不能为空")
	RegisterEmailEmpty    = xrrcode.New(100002, "注册邮箱号不能为空")
	RegisterUserNameEmpty = xrrcode.New(100003, "注册用户名不能为空")
	MobileHasRegistered   = xrrcode.New(100009, "手机号已经注册")
	LoginInfoEmpty        = xrrcode.New(100010, "登录参数错误")
	CodeKeyEmpty          = xrrcode.New(100011, "验证码发送失败")
	RePasswdEmpty         = xrrcode.New(100012, "密码修改参数错误，请重新输入")
)
