
# 压测简介
使用开源压测工具go-stress-testing，工具具体介绍如下
[go-stress-testing](https://github.com/link1st/go-stress-testing)


- 第一步下载 [go-stress-testing](https://github.com/link1st/go-stress-testing/releases)

- 第二步创建压测数据
```
curl 'http://127.0.0.1:8080/api/user/login' \
  -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
  -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8' \
  -H 'Cache-Control: max-age=0' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Cookie: _lxsdk_cuid=1761e65afb6c8-0a5332a8efff16-18346153-13c680-1761e65afb6c8; _lxsdk=1761e65afb6c8-0a5332a8efff16-18346153-13c680-1761e65afb6c8; Idea-d82114bc=95219ce2-c359-40a4-a167-5a9b4a296b11; moaDeviceId=1651CF34E26C52D29911A79D6FC97EFD; webNewUuid=8f427f110df783ffe6025357be63e783_1649322370760; JSESSIONID=DC6189C2C18E86727D7F5EDB10BF6710' \
  -H 'Origin: http://localhost:8080' \
  -H 'Referer: http://localhost:8080/api/index' \
  -H 'Sec-Fetch-Dest: document' \
  -H 'Sec-Fetch-Mode: navigate' \
  -H 'Sec-Fetch-Site: same-origin' \
  -H 'Sec-Fetch-User: ?1' \
  -H 'Upgrade-Insecure-Requests: 1' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36' \
  -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  --data-raw 'username=test4&password=1234567' \
  --compressed

```
- 第三步将上面请求数据写入curl/login.txt文件中

- 第四步了解指令参数意义
```
  -H value
    	自定义头信息传递给服务器 示例:-H 'Content-Type: application/json'
  -c uint
    	并发数 (default 1)
  -d string
    	调试模式 (default "false")
  -data string
    	HTTP POST方式传送数据
  -n uint
    	请求数(单个并发/协程) (default 1)
  -p string
    	curl文件路径
  -u string
    	压测地址
  -v string
    	验证方法 http 支持:statusCode、json webSocket支持:json
```


- 第五步执行指令
```
./go-stress-testing-mac -c 200 -n 100 -p curl/login.txt
```