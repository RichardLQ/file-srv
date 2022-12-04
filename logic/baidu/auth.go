package baidu

import (
	"fmt"
	"github.com/RichardLQ/file-srv/auth"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
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


