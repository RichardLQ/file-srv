package uploads

import (
	"github.com/RichardLQ/file-srv/auth"
	"github.com/RichardLQ/file-srv/refer"
)

//添加数据
func (u *UploadFile) Create() error {
	err := auth.Global.Mini.Table(refer.Table_Uploads).
		Create(u).Error
	if err != nil {
		return err
	}
	return err
}