package code

import "MyRainFarm/pkg/xrrcode"

var (
	RegisterEmpty       = xrrcode.New(110009, "注册失败")
	GetUserInfoEmpty    = xrrcode.New(110001, "获取用户信息失败")
	LoginInfoEmpty      = xrrcode.New(110002, "用户登录信息错误，登录失败")
	LoginEmpty          = xrrcode.New(110003, "用户登录失败")
	UserCodeEmpty       = xrrcode.New(110004, "用户验证码设置失败")
	UserCodeAgainEmpty  = xrrcode.New(110005, "用户验证码已经发送")
	UserCodeAccessEmpty = xrrcode.New(110006, "用户验证码错误")
)
