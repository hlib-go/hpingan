package hpingan

import (
	"encoding/json"
	"testing"
	"time"
)

const (
	// 私钥为Base64编码的内容
	RSA_SIGN_PRIVATE_KEY = "MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDAYYN4LuDzDmVpDaJQfKMUpuGcZEUuHY0/9PmhUedE7FjeDi7Pv2bo3mt2irPcnoyk1Kf8gJrAR8WIeSrYS2PViG4F0ZES5peQzTea0LKmqicdf14CMFuABM3wixlvVSiNNU6b7tuFAyGK+bM5O0Nq0DVLIfcFdQ/ISra2p5jIpnZytWdzvFZtVhKbcHHcWTrZYwBDIhuRzsWID0pKbrZ8xfJtaGEAubBBzIz8wMBD0cPChkHoAK0Do0dvplRTLYEw9FAff9DbznkMuhYNM50UfrnXJbFarUqZFafqpeog1YfIOuJXTh5AGcfT70jiJoqHlFkWSMLRXfOzdNJuScFXAgMBAAECggEAVHA4gR6qU2RFNIfsb0FVXvIkYj1MdAI6hhDjJGBLjt/h3Y0CFOMgqyM8raoRacsZVhuUo+pagfy+gbkVChyE9CtOhzVFZLNbYQ4B0uy5DRfv1S2bOIC7GISDU4nqHwiUpO35081SAp6uMU6J0JvnalF6osCEqJuJ0tKgQalsENzi5kfC407lKwZQneC45aenHU5yEu8BOLGD1wmskshNL/CZDsPE1IomA/evZMJFPbQRnjVNTQrm/iThtMKjcgs86TLszZAv2uSBAOunWO0aFrD9UGsrHi+fiOdAO/Aldxg3oDMdBkcKi0oDS0nSvPjHEwSqUzr696eeSYtCb34SSQKBgQD9bqyn7ph9L5QCRiUy/z5Lbnj64ZDmUZR5ISfAR0QZo9YKc/3CH/NTFNPa3nH8o0PON5r1xiVDrOqeoHigPar1HBUn7FACvsoWCWGo05Trg+iHw7arXg63iVdvq1Y++cJw60awhf1Oyc0BaU6GjZqndlsfN168bnCQoZDZVqmk5QKBgQDCVH2Ss1cf3HHD1GDx7mkF9gCtINbGjx33R8A7iiP3AhdeEfRbJgYN4NZQIPHp3pLlR8oURAe2wcJaPvJk3JcbHRr7F4bLx+o7QRqACK0qFiBGRAWTLW1+vSEF5B8XI7abJ0uh1/jY6SC5qRpULjFphFWqnPbyEBDaUxEsv5XFiwKBgQDb7LcX5A2gVTFz57tQTf+ZJf0Gechslk7p7sO4MCAAwBR2bcHAPEnDx0cxn+f6/RNSWR64OYMM/3l6vnbvV5SLsBsL0h58w2RLS0jcYP/SgV3Asy5J+A19aEngNidZ7xO1OfbWKPBw9t5YYzIpwah7ulHI/RvRGWXFnO8+K957oQKBgQCY2ZQbu2wf52ikO3w1TuzUf3Zwto+4hgFv6sPNEQ7QzphKqZylxPamG972D7O33BV2uky8O/ZFTYntKe3GX6uK7A24yfyAdLJMAR0y76AkplqkagRoiybAtUowNAowTSroRrQ6jIFzO4s9/UB0ThEXibtIA0ZJh1F6KDQuGlNXmQKBgQCzoKqgysE9zrN6za8hqvpiF7bbpWxYxHIcJFe5vS95Cm9IyoBXyi/jvSPtljIIfrNn1EysX1BKrw5jCCzLXT9rBUS901lUOy0xJQgkaLUwfg+JxcOf+sSqkjDnoAgP10rZ3/YrcMcaeZiWbkkO7h8a9ART3E8eVsEE0QqLaX5M8w=="
)

// 签名测试
func TestPinganRsaSign(t *testing.T) {
	sign, err := PinganRsaSign("123", RSA_SIGN_PRIVATE_KEY)
	if err != nil {
		t.Error(err)
	}
	t.Log(sign)
}

var c = &Pingan{
	Config: &Config{
		BaseServiceUrl:    "https://rsb-stg.pingan.com.cn",
		Channel:           "2134",
		ChannelPrivateKey: "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCvZ0HYdoCppkvmZ7lCgWYKmtiUIYEFiD6H1a8vrWS9iK0vRBbHKlxlBPqWhx+nDeO8e+fL3dMlmM6XkkhIQptTIpInvOTk1CigSfD/QmNJQ1v0O7+8RTzwaiCcg2+aknlcUJY7q6k43pWQQX7YiMJRwKO1ioM1deGG7Aa5sus8oIlnOEu0DPR398WFLHlbX5GSskWKjxA779dCG64USs6z4G2Jc7iDDaB5zi4t1HqW2cCoCQXy6ohiDzFjUBkAH1XmcCUMybqxmWn2MVi6cOx26/qyUMkY+8j+N0D1Cz04Qikw+qi+pPEH5Tkk33aZTqattO6yyNfG6+RjimgWW1VdAgMBAAECggEANbyesaBsfo0SB76TRUq7ZlhWLdut2XIKJjdS+npWuLJczHRBvja+I7zIwMEY82cP9DjJpO2KuKP99vW761dsmqWmvUHjpi9RWvoMOUevf0yBqWt8rENSDh2VbB3gXSyuvZd6z0EiEOgwTgeiwtfzIHSyMvYCiKhataktzkqAQwwxWBxhOAx2+9HOyub+FU4UkVxy55YOCaMs8plc/v1AeOBI4Iex3//ae4Mvf53bePHT5g36hQB0KwY23bzFaZbb+DrWDtnOmhVJFk+HG3AxYWa2bM5D/KroW0KpGsV35g0i+A1RDVZBJVXRM3KpASbYCmZLf1KuAwXURVj9nl/5SQKBgQDph8UatDF+qVf4f9eGjAZqUJqWHHX90tSq3asFyZh+x/tPvEvh73Ax53HpixxQuVemN17xMH3nD14AA3pjfwQ7auAmPeqnnmi9PAueGp7u8GSpN/npgyqyJcaRTi4vtq0n8evmbSinJHjEHE2r0Qrv0guySm4hvotjY0tPt4EgEwKBgQDAR7mocESI50Xt20I38CpWAf8mfd7b5PJj/IV4k2U7HIPPOUvQddv6ht0pFbWmjX8VXfLTq2kSF9T/PasvV2It2WTr1vTvekPJAyxQyOKUrCQFd4aGF4bRju9EJhQTC8GeKVphLaqD1LKTq0PHuYBqd9J4vy5zWkLhm4dZ90zCzwKBgFZ8WKvWFgpxjsnSFrCWhP20kjuLmMsC3I3XADJpgZdaWqEh+/xVn3hr3Yz/XuIqceJ12Sx7p/T9hPN+VkIo1RloMFEZVyDykpmadjb8ZuXYk37z4xLu584IHWzMxmAXSDNl7dZtkMOtxyfhz4YVKfa9FGXYGMdYeOi66jZAoIqPAoGAcXOgsycS+Shk1XSSvMjDYh4pe3+3WE3dS2u7ISrvBxLLHyEyf5ybIFYk9lHt7Yj9nBDj70rrYxaWwceUUwmfT2g/KPybZxXgGyH5Ku1s68YqNuOQh1woW0uYF8DUBWzhYD4MKty0CLceaw0ZoCXvBGMQ19gOQ1fVsKKBGW9WZncCgYEAvQWZ07pl5KcCmagz3tuSbH3VAu5G2caLrnUhaaU0f8eWkuMxJ2HVUnlf51EPQtVP2GOhKwo/iKf/eR9Maafu4+C+Yrh+m/R0P1tX//h34UhOGgrYzyvokW3q3XB/6i1oodyO6zdJ1jfcltI8jYnrEDLK/Z+LmIxxYHysuTnftwA=",
	},
	AccessToken: func(c *Pingan) string {
		// 根据渠道编号从表中查询，判断是否过期，如果过期重新获取
		// 650602：accessToken令牌已过期，请重新申请     650608：无效令牌，请重新申请

		t, _ := c.ApplyAccessToken()
		return t.AccessToken
	},
}

func TestClient_ApplyAccessToken(t *testing.T) {
	token, err := c.ApplyAccessToken()
	if err != nil {
		t.Error(err)
		return
	}
	b, _ := json.Marshal(&token)
	t.Log(string(b))
}

func TestPingan_ThreeAccountQuery(t *testing.T) {
	err := c.ThreeAccountQuery("13427125164", "111111111111111111", "20200605", "182520")
	if err != nil {
		t.Error(err)
	}
}

func TestPingan_ThreeFrozenFund(t *testing.T) {
	//err := c.ThreeFrozenFund("13427125164", 1, "211111111111111111", "20200605", "182520", "111111111111111111", "20200605", "182520")
	//if err != nil {
	//	t.Error(err)
	//}
}

func TestPingan_ThreeFrozenQuery(t *testing.T) {
	err := c.ThreeAccountQuery("13427125164", "111111111111111111", "20200605", "182520")
	if err != nil {
		t.Error(err)
	}
}

func TestPingan_Biz(t *testing.T) {
	// 13086672356
	// 可以交易成功的手机号：18038706921  15845522123
	trade, err := c.Biz("18038706921", 1)
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(trade)
	t.Log(string(bytes))
}

func TestTime(t *testing.T) {

	t.Log(time.Unix(1592558310150/1000, 0).Format(time.RFC3339))

}
