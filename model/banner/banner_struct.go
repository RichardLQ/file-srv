package banner

//Banners banner结构体
type Banners struct {
	Id int32 `json:"id"`
	Imageurl string `json:"imageurl"`
	Title string `json:"title"`
	Createtime int64 `json:"createtime"`
	Status int32 `json:"status"`
}

//Search 查询
type Search struct {
	Limit int32  `json:"limit"`
	Status int32 `json:"status"`
}