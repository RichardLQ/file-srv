package stub

import (
	"context"
	"fmt"
	"github.com/RichardLQ/fix-srv/logic"
	"github.com/RichardLQ/fix-srv/refer"
	"github.com/tidwall/gjson"
)

var UserService = &categoryServiceServer{}

type categoryServiceServer struct {
	FixServiceServer
}

//GetPicCategory 获取分类列表
func (this *categoryServiceServer) GetPicCategory1(ctx context.Context, req *GetPicCategory1Req) (*GetPicCategory1Rsp, error) {
	rsp := GetPicCategory1Rsp{
		Code: 200,
		Msg: "请求成功！",
	}
	data,err :=logic.GetPicCategory()
	if err != nil {
		rsp.Code = refer.Find_Piclist_Err
		rsp.Msg = "请求失败！"
	}
	for i, datum := range data {
		temp := &Photo{
			Id: int32(i+1),
			Pid: gjson.Get(datum.String(),"id").String(),
			Ename: gjson.Get(datum.String(),"ename").String(),
			Address: gjson.Get(datum.String(),"cover").String(),
			Zname: gjson.Get(datum.String(),"name").String(),
		}
		rsp.List = append(rsp.List,temp)
	}
	return &rsp,nil
}

//GetPicList1 获取分类列表
func (this *categoryServiceServer) GetPicList1(ctx context.Context, req *GetPicList1Req) (*GetPicList1Rsp, error) {
	rsp := GetPicList1Rsp{
		Code: 200,
		Msg: "请求成功！",
	}
	data,err :=logic.GetPicList1(req.Limit,req.Skip,req.Typeid)
	if err != nil {
		rsp.Code = refer.Find_Piccontent_Err
		rsp.Msg = "请求失败！"
	}
	for i, datum := range data {
		tem := []*Tags{}
		for i2, result := range gjson.Get(datum.String(),"tag").Array() {
			fmt.Println(i2)
			fmt.Println(result)
			t:= Tags{
				Id: int32(i2+1),
				Tag: result.String(),
			}
			tem = append(tem,&t)
		}
		temp := PicList1{
			Pid: int32(i+1),
			Imageid: gjson.Get(datum.String(),"id").String(),
			ImageSmall: gjson.Get(datum.String(),"thumb").String(),
			ImageMiddle: gjson.Get(datum.String(),"wp").String(),
			ImageBig: gjson.Get(datum.String(),"img").String(),
			Desc: gjson.Get(datum.String(),"desc").String(),
			Tags: tem,
		}
		rsp.List = append(rsp.List,&temp)
	}
	return &rsp,nil
}


//GetHotPicList1 热门分类图片内容
func (this *categoryServiceServer) GetHotPicList1(ctx context.Context, req *GetHotPicList1Req) (*GetPicList1Rsp, error) {
	rsp := GetPicList1Rsp{
		Code: 200,
		Msg: "请求成功！",
	}
	data,err :=logic.GetHotPicList1(req.Limit,req.Skip)
	if err != nil {
		rsp.Code = refer.Find_Piccontent_Err
		rsp.Msg = "请求失败！"
	}
	for i, datum := range data {
		tem := []*Tags{}
		for i2, result := range gjson.Get(datum.String(),"tag").Array() {
			fmt.Println(i2)
			fmt.Println(result)
			t:= Tags{
				Id: int32(i2+1),
				Tag: result.String(),
			}
			tem = append(tem,&t)
		}
		temp := PicList1{
			Pid: int32(i+1),
			Imageid: gjson.Get(datum.String(),"id").String(),
			ImageSmall: gjson.Get(datum.String(),"thumb").String(),
			ImageMiddle: gjson.Get(datum.String(),"wp").String(),
			ImageBig: gjson.Get(datum.String(),"img").String(),
			Desc: gjson.Get(datum.String(),"desc").String(),
			Tags: tem,
		}
		rsp.List = append(rsp.List,&temp)
	}
	return &rsp,nil
}






