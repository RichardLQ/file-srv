package banner

import (
	"fmt"
	"github.com/RichardLQ/fix-srv/client"
	"github.com/RichardLQ/fix-srv/refer"
)


//Find 查询
func (u *Search) Find() (*[]Banners ,error) {
	list := &[]Banners{}
	sql := client.Global.Mini.Table(refer.Table_Banner)
	if u.Limit != 0 {
		fmt.Println(333)
		sql = sql.Limit(u.Limit)
	}
	if u.Status != 0 {
		sql = sql.Where("status = ?",u.Status)
	}
	err:=sql.Find(list).Debug().Error
	if err !=nil {
		return list,err
	}
	return list,nil
}
