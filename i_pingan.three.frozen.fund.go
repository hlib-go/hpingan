package hpingan

import "encoding/json"

//平安三类户资金冻结
//说明：资金冻结，是指从当前交易发起时间到第二天23：59：59冻结资金，例如当前是2020.06.01, 冻结交易成功后，是从交易开始发起时一直冻结到2020.06.02 23：59：59。
func (c *Pingan) ThreeFrozenFund(mobile string, days int64, tradeId, submitDate, submitTime, oriTradeId, oriSubmitDate, oriSubmitTime string) (r *ThreeFrozenFundBankResponse, err *ErrCdoe) {
	bm := NewBodyMap()
	bm.Set("trxTp", "CNCPS316")
	bm.Set("mobile", mobile)
	bm.Set("days", days)
	bm.Set("tradeId", tradeId)       //客户端唯一交易流水号，18位，建议yyyymmdd+10位，支持数字和字母，字母区分大小写。例如202006031234567890
	bm.Set("submitDate", submitDate) //yyyyMMdd
	bm.Set("submitTime", submitTime) //HHmmss

	bm.Set("oriTradeId", oriTradeId)       //原查询交易
	bm.Set("oriSubmitDate", oriSubmitDate) //原查询交易
	bm.Set("oriSubmitTime", oriSubmitTime) //原查询交易

	data, err := c.Call(bm)
	if err != nil {
		return
	}
	if data != nil {
		bytes, e := json.Marshal(data)
		if e != nil {
			err = NewErr("ERR_JSON", e.Error())
			return
		}
		e = json.Unmarshal(bytes, &r)
		if e != nil {
			err = NewErr("ERR_JSON", e.Error())
			return
		}
	}
	return
}

type ThreeFrozenFundBankResponse struct {
	BankResponseCode string `json:"bankResponseCode"`
	BankResponseMsg  string `json:"bankResponseMsg"`
}
