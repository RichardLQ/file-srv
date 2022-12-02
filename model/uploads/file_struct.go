package uploads



//上传信息
type Uploads struct {
	Id         string `json:"id"`
	Uid        int `json:"uid"`
	Openid     string `json:"openid"`
	Address    string `json:"address"`
	Type       int    `json:"type"`
	Createtime string `json:"createtime"`
}

type UploadFile struct {
	Id         string `json:"id"`
	Uid 	   int `json:"uid"`
	Openid     string `json:"openid"`
	Address    string `json:"address"`
	Type       int    `json:"type"`
	Createtime string `json:"createtime"`
}