package system

import (
	"galen-gvm/global"
	"galen-gvm/model/response"
	"galen-gvm/model/system"
	"galen-gvm/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

func ApiRouter(c *gin.RouterGroup) {
	V1Group := c.Group("v1")
	{
		V1Group.GET("/getall", GetAllApis)
		V1Group.GET("/getApiById", GetApiById)
	}
}

// GetAllApis
// @Tags      SysApi
// @Summary   获取所有api
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  utils.Response{data=response.SysAPIListResponse}  "获取所有api"
// @Router    /system/api/v1/getall [get]
func GetAllApis(c *gin.Context) {
	apis, err := system.GetAllApis()
	if err != nil {
		global.GVA_LOG.Error("GetAllApis err:", zap.Error(err))
		utils.HTTPFail(nil, global.GetFail, c)
		return
	}
	utils.HTTPOk(response.SysAPIListResponse{Apis: apis}, global.GetSuccess, c)
}

// GetApiById
// @Tags      SysApi
// @Summary   根据id获取api
// @accept    application/json
// @Produce   application/json
// @Param     id  query      int                                   true  "根据id获取api"
// @Success   200   {object}  utils.Response{data=response.SysAPIResponse}  "根据id获取api,返回包括api详情"
// @Router    /system/api/v1/getApiById [get]
func GetApiById(c *gin.Context) {
	idQ := c.Query("id")
	id, err := cast.ToIntE(idQ)
	if err != nil || id < 1 {
		global.GVA_LOG.Error("GetApiById err:", zap.Error(err))
		utils.HTTPFail(nil, global.IDInvailed, c)
		return
	}
	api, err := system.GetApiById(id)
	if err != nil {
		global.GVA_LOG.Error("GetApiById db err:", zap.Error(err))
		utils.HTTPFail(nil, global.GetFail, c)
		return
	}
	utils.HTTPOk(response.SysAPIResponse{Api: api}, global.GetSuccess, c)
}
