package hpingan

import "encoding/json"

//平安三类户冻结交易结果查询
func (c *Pingan) ThreeFrozenQuery(mobile, tradeId, submitDate, submitTime, oriTradeId, oriSubmitDate, oriSubmitTime string) (r *ThreeFrozenQueryResult, err *ErrCdoe) {
	bm := NewBodyMap()
	bm.Set("trxTp", "CNCPS317")
	bm.Set("mobile", mobile)
	bm.Set("tradeId", tradeId)
	bm.Set("submitDate", submitDate) //yyyyMMdd
	bm.Set("submitTime", submitTime) //HHmmss

	bm.Set("oriTradeId", oriTradeId)       //原冻结交易
	bm.Set("oriSubmitDate", oriSubmitDate) //原冻结交易
	bm.Set("oriSubmitTime", oriSubmitTime) //原冻结交易

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
		err = NewErr("ERR_JSON", e.Error())
	}
	return
}

type ThreeFrozenQueryResult struct {
	OriResponseCode string `json:"oriresponseCode" description:"原冻结交易响应码"`
	OriResponseMsg  string `json:"oriresponseMsg" description:"原冻结交易响应码说明"`
}
