package auth

import (
	"github.com/jinzhu/gorm"
)

var Global = struct {
	Mini            gorm.DB   `confs:"name=Mini;read=rainbow"`
	FileServiceConf FileConf  `confs:"read=rainbow;format=json"`
	Qiniu_OSS       QiniuConf `confs:"read=rainbow;format=json"`
	BaiduKey        BaiduConf `confs:"read=rainbow;format=json"`
	//Temporarily  redis.Conn `confs:"name=Temporarily;read=rainbow"`
}{}

//BaiduConf 运行端口
type BaiduConf struct {
	AppID     string
	Appkey    string
	Secretkey string
	Signkey   string
	Code      string
}

//QiniuConf 运行端口
type QiniuConf struct {
	ACCESS_KEY string
	SECRET_KEY string
}

//FileConf 运行端口
type FileConf struct {
	HttpPort string
	HttpIp   string
}
