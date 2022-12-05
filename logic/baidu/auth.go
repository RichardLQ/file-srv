package baidu

import (
	"fmt"
	"github.com/RichardLQ/file-srv/auth"
	"github.com/RichardLQ/file-srv/util/baidu"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"mime/multipart"
	"net/http"
)
//获取code
//https://openapi.baidu.com/oauth/2.0/authorize?response_type=code&client_id=0WGiQErUYXHFRN2GWU5rzXc5NhMUGNg5&redirect_uri=
//oob&scope=basic,netdisk&display=tv&qrcode=1&force_login=1


//GetBaiduAccessToken 获取accesstoken
func GetBaiduAccessToken(c *gin.Context)  {
	token,err := setAccessToken()
	if err!= nil {
		fmt.Println(err)
		token,err = getAccessToken()
		if err!=nil {
			c.JSON(http.StatusOK, gin.H{"ret":"失败","token":token,"err":err,"code":1001})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"ret":"成功","token":token,"err":err,"code":http.StatusOK})
	return
}

//GetBaiduPrecreate 预上传
func GetBaiduPrecreate(c *gin.Context)  {
	token,_:=c.GetPostForm("token")
	files,_:=c.FormFile("files")
	path := "/uploads/" + files.Filename
	size := files.Size
	fmt.Println(size)
	var blockList []string
	listString,_:=auth.FileOpenMD5(files)
	blockList = append(blockList, listString)
	arg := baidu.NewPrecreateArg(path, uint64(size), blockList)
	if ret, err := baidu.Precreate(token, arg); err != nil {
		fmt.Printf("[msg: precreate error] [err:%v]", err.Error())
	} else {
		getBaiduSuperfile(ret.UploadId,path,token,files)
		getBaiduCreate(ret.UploadId,path,token,blockList,uint64(size))
	}
}

//getBaiduCreate 分片上传
func getBaiduCreate(uploadId,path,token string,blockList []string,size uint64)  {
	// // call create API
	arg := baidu.NewCreateArg(uploadId, path, size, blockList)
	if ret, err := baidu.Create(token, arg); err != nil {
		fmt.Printf("[msg: create this part error] [err:%v]", err.Error())
	} else {
		fmt.Printf("ret:%+v\n", ret)
	}
}


//GetBaiduSuperfile2 分片上传
func getBaiduSuperfile(uploadId,path,token string,file *multipart.FileHeader)  {
	partseq := 0
	arg := baidu.NewUploadArg(uploadId, path, file, partseq)
	if ret, err := baidu.Upload(token, arg); err != nil {
		fmt.Printf("[msg: upload this part error] [err:%v]", err.Error())
	} else {
		fmt.Printf("ret:%+v", ret)
	}
}

func getAccessToken() (string,error)  {
	conn := auth.RedisConn()
	value,err:=redis.String(conn.Do("GET","refresh_token"))
	if err != nil {
		return "",err
	}
	url := "https://openapi.baidu.com/oauth/2.0/token?grant_type=refresh_token&refresh_token=" +
		value+"&client_id="+auth.Global.BaiduKey.Appkey+"&client_secret="+auth.Global.BaiduKey.Secretkey
	listString, err := auth.SendHttpRequest(url, "", "", nil, nil)
	if err!= nil {
		return "",err
	}
	token:=gjson.Get(listString,"access_token").String()
	if token == "" {
		return "",fmt.Errorf("code过期")
	}
	conn.Do("set","access_token",gjson.Get(listString,"access_token").String())
	conn.Do("set","refresh_token",gjson.Get(listString,"refresh_token").String())
	return token, nil
}

//setAccessToken 获取token
func setAccessToken() (string,error) {
	url := "https://openapi.baidu.com/oauth/2.0/token?grant_type=authorization_code&code="+auth.Global.BaiduKey.Code+"&client_id=" +
		auth.Global.BaiduKey.Appkey +"&client_secret="+ auth.Global.BaiduKey.Secretkey +"&redirect_uri=oob"
	listString, err := auth.SendHttpRequest(url, "", "", nil, nil)
	if err!= nil {
		return "",err
	}
	token:=gjson.Get(listString,"access_token").String()
	if token == "" {
		return "",fmt.Errorf("code过期")
	}
	conn := auth.RedisConn()
	conn.Do("set","access_token",gjson.Get(listString,"access_token").String())
	conn.Do("set","refresh_token",gjson.Get(listString,"refresh_token").String())
	return token,nil
}


