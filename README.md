# 平安银行网关接入SDK

fat:

https://rsb-stg.pingan.com.cn/brop/front/openapi/restful/ ch-分配给你的渠道号
    /brcp-front-proxy.applyAccessToken?channel=分配给你的渠道号&random=生成的随机
    数&timestamp=当前时间戳&sign=签名结果
    
PRD： 

https://rsb.pingan.com.cn/brop/front/openapi/restful/ ch-分配给你的渠道号
/brcp-front-proxy.applyAccessToken?channel=分配给你的渠道号&random=生成的随机
数&timestamp=当前时间戳&sign=签名结果     

获取到：accessToken 与过期时间 

==》服务请求
https://rsb-stg.pingan.com.cn/brop/front/openapi/restful/ch-分配给你的渠道号/{serviceName}
body:{"channel":"xxxxxx","accessToken":"xxxxxx",其他业务参数}  



