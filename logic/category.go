package logic

import (
	"fmt"
	"github.com/RichardLQ/fix-srv/auth"
	"github.com/RichardLQ/fix-srv/client"
	"github.com/RichardLQ/fix-srv/refer"
	"github.com/tidwall/gjson"
)

//GetPicCategory 获取图片分类
func GetPicCategory() ([]gjson.Result,error) {
	url := refer.API_HOST + "v1/vertical/category"
	listString, err := auth.SendHttpRequest(url, "", "", nil, nil)
	if err != nil {
		return []gjson.Result{},err
	}
	data := gjson.Get(listString, "res.category").Array()

	return data,nil
}

//GetPicList1 分类图片内容
func GetPicList1(limit,skip int32,typeid string) ([]gjson.Result,error) {
	if client.Global.SelectType.Type == "2" {
		typeid = "5109e04e48d5b9364ae9ac45"
	}
	url := fmt.Sprintf(refer.API_HOST+"v1/vertical/category/%s/vertical?adult=false" +
		"&first=1&order=hot&limit=%d&skip=%d",typeid,limit,skip)
	listString, err := auth.SendHttpRequest(url, "", "", nil, nil)
	if err != nil {
		return []gjson.Result{},err
	}
	data := gjson.Get(listString, "res.vertical").Array()
	return data,nil
}

//GetHotPicList1 热门分类图片内容
func GetHotPicList1(limit,skip int32) ([]gjson.Result,error) {
	if client.Global.SelectType.Type == "2" {
		skip = 1
	}
	url := fmt.Sprintf(refer.API_HOST+"v1/vertical/vertical?adult=false&first=20&order=new&limit=%d&skip=%d",limit,skip)
	listString, err := auth.SendHttpRequest(url, "", "", nil, nil)
	if err != nil {
		return []gjson.Result{},err
	}
	data := gjson.Get(listString, "res.vertical").Array()
	return data,nil
}



