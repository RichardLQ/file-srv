package uploads

import (
	"github.com/RichardLQ/file-srv/refer"
	"github.com/RichardLQ/fix-srv/client"
)

//添加数据
func (u *UploadFile) Create() error {
	err := client.Global.Mini.Table(refer.Table_Uploads).
		Create(u).Error
	if err != nil {
		return err
	}
	return err
}