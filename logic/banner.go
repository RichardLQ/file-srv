package logic

import "github.com/RichardLQ/fix-srv/model/banner"

//GetBanner banner列表
func GetBanner(limit ,status int32) (*[]banner.Banners,error) {
	ban :=banner.Search{
		Limit: limit,
		Status: status,
	}
	list,err := ban.Find()
	if err!= nil {
		return list,err
	}
	return list,nil
}