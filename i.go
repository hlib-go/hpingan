package hpingan

import (
	"crypto"
	"crypto/md5"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/hlib-go/hgenid"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	YYYYMMDDHHMMSS = "20060102150405"
	YYYYMMDD       = "20060102"
	DDHHMMSS       = "150405"
)

// 银联云闪付接口请求客户端
type Pingan struct {
	Config *Config

	// 客户端实现读取token逻辑，避免每次调用接口获取token
	AccessToken func(c *Pingan) string
}

func New(config *Config, accessToken func(c *Pingan) string) *Pingan {
	return &Pingan{
		Config:      config,
		AccessToken: accessToken,
	}
}

// 返回 ERR_HTTP 错误时，需要调用查询接口验证冻结是否成功
func (c *Pingan) Call(bm *BodyMap) (result interface{}, err *ErrCdoe) {
	plog := log.WithField("requestId", Rand32()).WithField("cname", "pingan").WithField("channel", c.Config.Channel)
	bm.Set("channel", c.Config.Channel)
	bm.Set("accessToken", c.AccessToken(c))
	body, e := json.Marshal(bm.m)
	if e != nil {
		plog.Error(e)
		err = ERR_JSON
		return
	}
	url := c.Config.BaseServiceUrl + "/brop/front/openapi/restful/ch-" + c.Config.Channel + "/http_bcard-cncps-91732.pab.com.cn:40001_cncps-acct_cncpsservlet_pafaOpenApiRequest"
	plog.Info("平安请求地址:" + url)
	plog.Info("平安请求报文:" + string(body))
	resp, e := http.Post(url, "application/json", strings.NewReader(string(body)))
	if e != nil {
		plog.Error(e)
		err = ERR_HTTP
		return
	}
	if resp.StatusCode != 200 {
		e = errors.New(resp.Status)
		plog.Error(e)
		err = ERR_HTTP
		return
	}
	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		plog.Error(e)
		err = ERR_HTTP
		return
	}
	plog.Info("平安响应报文:" + string(bytes))

	// 解析响应报文
	var rmap map[string]interface{}
	e = json.Unmarshal(bytes, &rmap)
	if e != nil {
		plog.Error(e)
		err = ERR_HTTP
		return
	}
	respCode := rmap["responseCode"].(string) // 网关响应码
	respMsg := rmap["responseMsg"].(string)   // 网关响应码说明
	if respCode == E650602.Code {
		// accessToken令牌已过期，请重新申请
		//return c.Call(bm,result)
	}
	if respCode != E000000.Code {
		err = NewErr(respCode, respMsg)
		return
	}

	// 解析报文data
	e = json.Unmarshal([]byte(rmap["data"].(string)), &rmap)
	if e != nil {
		plog.Error(e)
		err = ERR_HTTP
		return
	}
	bankResponseCode := rmap["bankResponseCode"].(string) // 银行响应码
	bankResponseMsg := rmap["bankResponseMsg"].(string)   // 银行响应码说明
	if bankResponseCode != E0000.Code {
		err = NewErr(bankResponseCode, bankResponseMsg)
		return
	}
	result = rmap
	return
}

func (c *Pingan) ApplyAccessToken() (token *ApplyAccessToken, err error) {
	plog := log.WithField("requestId", Rand32()).WithField("cname", "pingan").WithField("channel", c.Config.Channel)
	random := GetRandomString(16)
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)

	// 由channelId、random和timestamp生成的数字签名。 //data = channel + ";" + random + ";" + times;
	signString := c.Config.Channel + ";" + random + ";" + timestamp
	sign, err := PinganRsaSign(signString, c.Config.ChannelPrivateKey)
	if err != nil {
		return
	}
	url := c.Config.BaseServiceUrl + "/brop/front/openapi/restful/ch-" + c.Config.Channel + "/brcp-front-proxy.applyAccessToken?channel=" + c.Config.Channel + "&random=" + random + "&timestamp=" + timestamp + "&sign=" + sign
	plog.Info("请求URL " + url)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	plog.Info("获取token响应报文：" + string(bytes))
	err = json.Unmarshal(bytes, &token)
	if err != nil {
		return
	}
	if token.ResponseCode != "0" {
		return nil, errors.New(token.ResponseCode + ":" + token.ResponseMsg)
	}
	return
}

type ApplyAccessToken struct {
	ResponseCode string `json:"responseCode"`
	ResponseMsg  string `json:"responseMsg"`
	AccessToken  string `json:"accessToken"`
	ExpireTime   int64  `json:"expireTime"`
}

// 获取随机字符串
//    length：字符串长度
func GetRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	var (
		result []byte
		b      []byte
		r      *rand.Rand
	)
	b = []byte(str)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

// UUID
func UUID() string {
	return Rand32()
}

// 平安RSA签名
func PinganRsaSign(v string, channelPrivateKey string) (sign string, err error) {
	priKey, err := base64.StdEncoding.DecodeString(channelPrivateKey)
	if err != nil {
		return
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(priKey)
	if err != nil {
		return
	}
	hash := md5.New()
	hash.Write([]byte(v))
	bytes := hash.Sum(nil)

	b, err := rsa.SignPKCS1v15(cryptorand.Reader, privateKey.(*rsa.PrivateKey), crypto.MD5, bytes)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(b)
	return
}

func genTradeId() string {
	return time.Now().Format(YYYYMMDD) + hgenid.GenId()[9:]
}

type TradeResult struct {
	TradeId  string
	RespCode string
	RespMsg  string
	Days     int64
	Mobile   string
}

func (c *Pingan) Biz(mobile string, days int64) (trade *TradeResult, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Error(e)
			err = errors.New("99999:panic 调用平安接口异常")
		}
	}()

	// 查询账户
	tradeId := genTradeId() // 格式：8位日期+10位数字
	submitDate := time.Now().Format(YYYYMMDD)
	submitTime := time.Now().Format(DDHHMMSS)
	e := c.ThreeAccountQuery(mobile, tradeId, submitDate, submitTime)
	if e != nil {
		err = errors.New(e.Error())
		return
	}

	// 冻结资金
	tradeId1 := genTradeId()
	submitDate1 := time.Now().Format(YYYYMMDD)
	submitTime1 := time.Now().Format(DDHHMMSS)
	r, e := c.ThreeFrozenFund(mobile, days, tradeId1, submitDate1, submitTime1, tradeId, submitDate, submitTime)

	trade = &TradeResult{
		TradeId: tradeId1,
		Days:    days,
		Mobile:  mobile,
	}
	if e == nil {
		if r != nil {
			trade.RespCode = r.BankResponseCode
			trade.RespMsg = r.BankResponseMsg
		}
		return
	}

	// HTTP请求错误，执行查询
	if e == ERR_HTTP {
		// 冻结查询
		tradeId2 := genTradeId()
		submitDate2 := time.Now().Format(YYYYMMDD)
		submitTime2 := time.Now().Format(DDHHMMSS)
		r, e := c.ThreeFrozenQuery(mobile, tradeId2, submitDate2, submitTime2, tradeId1, submitDate1, submitTime1)
		if e != nil {
			return trade, errors.New(e.Error())
		}
		trade.RespCode = r.OriResponseCode
		trade.RespMsg = r.OriResponseMsg
	} else if e != nil {
		err = errors.New(e.Error())
		return
	}

	if trade.RespCode != E0000.Code {
		err = errors.New(trade.RespCode + ":" + trade.RespMsg)
	}
	return trade, err
}
