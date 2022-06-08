package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
)

// ----社区相关

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("Logic.GetCommunityList() failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)  // 不轻易把服务器端报错暴露给外面【前端】
		return
	}
	ResponseSuccess(c,data)
	return
}
// CommunityDetailHandler 获取社区详情
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("id")
	id ,err := strconv.ParseInt(idStr,18,64)
	if err != nil{
		ResponseError(c,CodeInvalidParam)
		return
	}

	//
	data,err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("Logic.GetCommunityDetail() failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)  // 不轻易把服务器端报错暴露给外面【前端】
		return
	}
	ResponseSuccess(c,data)
	return
}
