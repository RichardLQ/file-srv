package route

import (
	"github.com/RichardLQ/file-srv/logic"
	"github.com/gin-gonic/gin"
)

func AuthRouter(e *gin.Engine) {
	v1 := e.Group("/auth")
	{
		v1.GET("/baiduAuth", logic.UploadImage)//百度获取accesstoken
	}
}