package auth

import "github.com/jinzhu/gorm"


var Global = struct {
	Mini            gorm.DB  `confs:"name=Mini;read=rainbow"`
	FileServiceConf FileConf `confs:"read=rainbow;format=json"`
}{}

//FileConf 运行端口
type FileConf struct {
	HttpPort string
	HttpIp   string
}