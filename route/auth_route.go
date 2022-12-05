package route

import (
	"github.com/RichardLQ/file-srv/logic/baidu"
	"github.com/gin-gonic/gin"
)

func AuthRouter(e *gin.Engine) {
	v1 := e.Group("/auth")
	{
		v1.GET("/baiduAuth", baidu.GetBaiduAccessToken)//百度获取accesstoken
		v1.POST("/getBaiduPrecreate", baidu.GetBaiduPrecreate)//百度上传
	}
}