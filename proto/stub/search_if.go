package stub

import (
	"context"
	"github.com/RichardLQ/fix-srv/logic"
	"github.com/tidwall/gjson"
)

//GetSearchList 热门分类图片内容
func (this *categoryServiceServer) GetSearchList(ctx context.Context, req *GetSearchReq) (*GetSearchRsp, error) {
	rsp := GetSearchRsp{
		Code: 200,
		Msg: "请求成功！",
	}
	data,_ :=logic.GetSearchList(req.Keyword,req.Start)
	for i, datum := range data {
		temp := SearchList{
			Sid: int32(i+1),
			Album: gjson.Get(datum.String(),"album.covers.0").String(),
			Photo: gjson.Get(datum.String(),"photo.path").String(),
			Sender: gjson.Get(datum.String(),"sender.avatar").String(),
			Dates: gjson.Get(datum.String(),"add_datetime").String(),
		}
		rsp.List = append(rsp.List,&temp)
	}
	return &rsp,nil
}





