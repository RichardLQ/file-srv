package uploads



//上传信息
type Uploads struct {
	Id         string `json:"id"`
	Openid     string `json:"openid"`
	Address    string `json:"address"`
	Type       int    `json:"type"`
	Createtime string `json:"createtime"`
}

type UploadFile struct {
	Id         string `json:"id"`
	Openid     string `json:"openid"`
	Address    string `json:"address"`
	Type       int    `json:"type"`
	Createtime string `json:"createtime"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
}