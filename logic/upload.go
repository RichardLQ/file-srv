package logic

import (
	"github.com/RichardLQ/file-srv/model/uploads"
	"github.com/RichardLQ/file-srv/refer"
	"github.com/RichardLQ/file-srv/util/qiniu"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//上传图片
func UploadImage(c *gin.Context)  {
	file, err := c.FormFile("file")
	openid := c.PostForm("openid")
	types,_:=strconv.Atoi(c.PostForm("types"))
	if openid == ""  {
		c.JSON(http.StatusOK, gin.H{"ret":"查询失败","data":"","err":"参数缺失","code":refer.PARAM_LACK})
		return
	}
	if err != nil {
		c.String(500, "上传图片出错")
	}
	str1,_,_,err:=qiniu.QiNiu_SourceUploadFile(file,"load","")
	if err !=nil {
		c.JSON(http.StatusOK, gin.H{"ret":"上传失败","data":"","err":err,"code":-1})
		return
	}
	if str1 != "" {
		pic := uploads.UploadFile{
			Openid: openid,
			Address: str1,
			Type: types,
			Createtime: time.Now().Format(refer.FORMATDATELONG),
		}
		err:=pic.Create()
		if err !=nil {
			c.JSON(http.StatusOK, gin.H{"ret":"上传失败","data":"","err":err,"code":-2})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"ret":"成功","data":str1,"err":"","code":http.StatusOK})
	return
}