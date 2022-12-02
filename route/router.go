package route

import (
	"github.com/RichardLQ/file-srv/logic"
	"github.com/gin-gonic/gin"
)

func IndexRouter(e *gin.Engine) {
	v1 := e.Group("/file")
	{
		v1.POST("/uploadImage", logic.UploadImage)//上传图片
	}
}