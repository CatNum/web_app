package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"web_app/models"
)

//把每一步数据库操作封装成函数
//待logic层根据业务需求调用

// 密码加密的盐
const secret = "yunshan"

var (
	ErrorUserExist = errors.New("用户已存在")
	ErrorUserNotExist = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

// CheckUserExist 检查指定用户名的用户是否存在，存在则返回错误
func CheckUserExist(username string) (err error) {
	var count int64
	result := DB.Where("username = ?", username).Find(&models.User{}).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count > 0 {
		return ErrorUserExist
	}
	return err
}

// CheckUserNotExist 检查指定用户名的用户是否存在，不存在则返回错误
func CheckUserNotExist(username string) (err error) {
	var count int64
	result := DB.Where("username = ?", username).Find(&models.User{}).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count == 0 {
		return ErrorUserNotExist
	}
	return err
}

// InsterUser 向数据库中插入一条新的用户记录
func InsterUser(user *models.User) (err error) {
	//对密码及进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	result := DB.Create(user)
	err = result.Error
	if err != nil {
		fmt.Println("err:", result.Error)
	}
	return err
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	//加盐
	h.Write([]byte(secret))
	//字节类型转换为16进制的字符串
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// InitUser 建立用户表
func InitUser() (err error) {
	user := models.User{}
	exist := DB.Migrator().HasTable(&user)
	if !exist{
		err = DB.AutoMigrate(&user)
	}
	return err
}

func Login(user *models.User) (err error) {
	oPassword := encryptPassword(user.Password)
	loginUser := new(models.ParamLogin)
	err = CheckUserNotExist(user.Username)
	if err != nil {
		return err
	}
	// 智能搜索，并将查询到的信息存到 loginUser 对象中
	//select username,password from users where username = user.Username;
	result := DB.Model(&models.User{}).Where("username", user.Username).Find(&loginUser)
	if result.Error != nil {
		return result.Error
	}
	if loginUser.Password != oPassword {
		return ErrorInvalidPassword
	}
	return err
}
