package hpingan

type ErrCdoe struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func NewErr(code, msg string) *ErrCdoe {
	return &ErrCdoe{
		Code: code,
		Msg:  msg,
	}
}

func (e *ErrCdoe) Error() string {
	return e.Code + ":" + e.Msg
}

var (
	E000000 = NewErr("000000", "网关响应成功")
	E0000   = NewErr("0000", "接口响应成功")
	E650602 = NewErr("650602", "accessToken令牌已过期，请重新申请")
	E0010   = NewErr("0010", "根据手机号未找到开卡记录")
	E1002   = NewErr("0002", "余额不足，不满足参与活动条件")

	ERR_JSON = NewErr("ERR_JSON", "转JSON出错")
	ERR_HTTP = NewErr("ERR_HTTP", "HTTP请求出错")
)

/*
0000	交易成功
0010	根据手机号未找到开卡记录
0011	手机号非首次开户，无法参与活动
0012	开卡日期不满足参与活动条件；
1001	根据卡号查询客户余额失败
1002	余额不足，不满足参与活动条件
1003	已冻结，无法继续冻结
1004	账户还未冻结
1005	余额不足，冻结失败
2001	未找到原交易（活动资格查询或资金冻结）
2002	冻结余额失败
2003	原交易失败，无活动资格
6196	传入的参数不全
8999	数据库异常
9001	交易超时

*/
