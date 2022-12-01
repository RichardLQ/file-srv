package refer

const (
	//登录失败
	Login_Miss = 1001//登录错误
	Token_err = 1002//token错误
	Params_err = 1003 //参数缺失
	Del_User_Err = 2001//用户删除错误
	Find_User_Err = 2002//查询用户错误
	Update_User_Err = 2003//更新用户错误

	Find_Collect_Err = 3001//查询收藏错误
	Update_Collect_Err = 3002//更新收藏错误

	Find_Piclist_Err = 4001 //查询分类列表错误
	Find_Piccontent_Err = 4002 //查询分类内容错误

	Find_Banner_Err = 4003 //查询banner错误
)
