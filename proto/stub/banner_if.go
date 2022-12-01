package stub

import (
	"context"
	"github.com/RichardLQ/fix-srv/logic"
	"github.com/RichardLQ/fix-srv/refer"
)

//GetBanner 获取banner栏
func (this *categoryServiceServer) GetBanner(ctx context.Context, req *GetBannerReq) (*GetBannerRsp, error) {
	rsp := GetBannerRsp{
		Code: 200,
		Msg: "请求成功！",
	}
	list,err :=logic.GetBanner(req.Limit,req.Status)
	if err != nil {
		rsp.Code = refer.Find_Banner_Err
		rsp.Msg = "请求失败！"
	}
	for i, datum := range *list {
		status := SourceTypes_Journey
		if datum.Status == 2 {
			status = SourceTypes_Memory
		}
		temp := BannerList{
			Bid: int32(i+1),
			Imageurl: datum.Imageurl,
			Title: datum.Title,
			Sourcetype: status,
		}
		rsp.List = append(rsp.List,&temp)
	}
	return &rsp,err
}
