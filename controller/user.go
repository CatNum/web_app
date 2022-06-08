package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"
)
// SignUpHandler 处理注册函数
func SignUpHandler(c *gin.Context){
	// 1.获取参数，参数校验
	 p := new(models.ParamSignUp)
	// ShouldBindJSON 只能判断字段类型（是否string）、格式（是否json）判断
	if err := c.ShouldBindJSON(&p);err != nil{
		//记录错误日志
		zap.L().Error("SignUp with invalid param",zap.Error(err))
		//判断 err 是否 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		//参数有误，返回响应
		if !ok {
			ResponseError(c,CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParam,removeTopStruct(errs.Translate(trans)))//翻译错误
		return
	}
	// 2.业务处理
	if err := logic.SignUp(p);err != nil{
		if errors.Is(err,mysql.ErrorUserExist){
			ResponseError(c,CodeUserExist)
			return
		}
		ResponseError(c,CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c,nil)
}
// LoginHandler 用户登录函数
func LoginHandler(c *gin.Context)  {
	// 1.获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p);err != nil {
		//记录错误日志
		zap.L().Error("Login with invalid param",zap.Error(err))
		//判断 err 是否 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		//参数有误，返回响应
		if !ok {
			ResponseError(c,CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParam,removeTopStruct(errs.Translate(trans))) //翻译错误
		return
	}
	// 2.业务逻辑处理
	token,err :=logic.Login(p)
	if err != nil{
		zap.L().Error("logic.Login failed",zap.String("username",p.Username),zap.Error(err))
		if errors.Is(err,mysql.ErrorUserNotExist){
			ResponseError(c,CodeUserNotExist)
			return
		}
		ResponseError(c,CodeInvalidPassword)
		return
	}
	// 3.返回响应
	ResponseSuccess(c,token)
}
