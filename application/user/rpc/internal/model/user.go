package model

import (
	"MyRainFarm/application/user/rpc/internal/code"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type User struct {
	SnowID   int64  `json:"snow" gorm:"column:snow"`
	Name     string `json:"name" gorm:"column:name"`
	Password string `json:"password" gorm:"column:password"`
	Mobile   string `json:"mobile" gorm:"column:mobile"`
	Email    string `json:"email" gorm:"column:email"`
}

func (m *User) TableName() string {
	return "user"
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db,
	}
}

func (m *UserModel) Insert(ctx context.Context, data *User) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *UserModel) GetUserInfo(ctx context.Context, snow string, data *User) error {
	return m.db.WithContext(ctx).Where("snow = ?", snow).First(data).Error
}

func (m *UserModel) GetUserPassword(ctx context.Context, mobile, password string) (int64, error) {
	data := &User{}
	err := m.db.WithContext(ctx).Where("mobile = ?", mobile).First(data).Error
	if err != nil {
		logx.Errorf("数据库查询失败", err)
		return 0, err
	}
	if password == data.Password {
		return data.SnowID, nil
	}
	return 0, code.LoginInfoEmpty
}

func (m *UserModel) GetUserInfoForEmail(ctx context.Context, email string, data *User) error {
	return m.db.WithContext(ctx).Where("email = ?", email).First(data).Error
}

func (m *UserModel) UpdateUserPwd(ctx context.Context, phoneNum, pwd string) (int64, error) {
	err := m.db.WithContext(ctx).Model(&User{}).Where("mobile = ?", phoneNum).Update("password = ?", pwd).Error
	if err != nil {
		logx.Errorf("更新密码失败", err)
		return 0, err
	}
	data := &User{}
	err = m.db.WithContext(ctx).Where("mobile = ?", phoneNum).First(data).Error
	if err != nil {
		logx.Errorf("数据库查询失败", err)
		return 0, err
	}
	return data.SnowID, nil
}
