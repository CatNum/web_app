package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
// 前后端分离
	{
		"code":1001, // 程序中的错误码
		"msg":"xx"   // 提示信息
		"data":{}    // 数据
	}
*/
// 定义一个结构体来实现这个功能
type ResponseData struct {
	Code ResCode
	Msg  interface{} // 因为 Msg 比较多变
	Data interface{} // 因为不知道 Data 的具体类型
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}
// 返回自定义msg
func ResponseErrorWithMsg(c *gin.Context, code ResCode,msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

