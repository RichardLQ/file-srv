package logic

import (
	"fmt"
	"github.com/RichardLQ/fix-srv/auth"
	"github.com/RichardLQ/fix-srv/client"
	"github.com/tidwall/gjson"
)

func GetSearchList(keyword string,start int32) ([]gjson.Result,error){
	if client.Global.SelectType.Type == "2" {
		keyword = "精美壁纸"
	}
	urls := "https://www.duitang.com/napi/blogv2/list/by_search/?kw="+keyword+"&after_id="+ fmt.Sprint(start) + "&type=feed&include_fields=top_comments%2Cis_root%2Csource_link%2Citem%2Cbuyable%2Croot_id%2Cstatus%2Clike_count%2Clike_id%2Csender%2Calbum%2Creply_count%2Cfavorite_blog_id&_type=&_=1639233452849"
	fmt.Println(urls)
	listString, err := auth.SendHttpRequest(urls, "", "", nil, nil)
	if err != nil {
		return []gjson.Result{},err
	}
	data := gjson.Get(listString, "data.object_list").Array()
	return data,nil
}