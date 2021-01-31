package hpingan

// 平安三类户账户活动资格查询

func (c *Pingan) ThreeAccountQuery(mobile, tradeId, submitDate, submitTime string) (err *ErrCdoe) {
	bm := NewBodyMap()
	bm.Set("trxTp", "CNCPS315")
	bm.Set("mobile", mobile)
	bm.Set("tradeId", tradeId)       //客户端唯一交易流水号，18位，建议yyyymmdd+10位，支持数字和字母，字母区分大小写。例如202006031234567890
	bm.Set("submitDate", submitDate) //yyyyMMdd
	bm.Set("submitTime", submitTime) //HHmmss

	_, err = c.Call(bm)
	return
}
