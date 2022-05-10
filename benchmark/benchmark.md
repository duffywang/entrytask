
# 压测服务器
- 机器个数：1
- CPU：8核
- 内存：16G
- 系统：MaxOS
- 数据库MySQL：5.6.14 
- 数据库数据量：10,000,000

# 观察指标
主要关注下面指标
- QPS
- 失败率
- 平均耗时
- 最长耗时
- TP90 TP99 TP999


# 压测工具简介
使用开源压测工具go-stress-testing，工具具体介绍如下
[go-stress-testing](https://github.com/link1st/go-stress-testing)。


- 第一步下载 [go-stress-testing](https://github.com/link1st/go-stress-testing/releases)。

- 第二步创建压测数据。
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
- 第三步将上面请求数据写入curl/login.txt文件中。

- 第四步选择压测参数。
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
  -m int 
      连接数
  -k 
     开启长连接         
  -u string
    	压测地址
  -v string
    	验证方法 http 支持:statusCode、json webSocket支持:json
```

# 压测结果
## 压测结论
1. 所有压测请求均返回正确的结果
2. 200并发固定用户(get接口)QPS达到11000，达到QPS大于3000的目标。
3. 200并发随机用户(login接口)QPS达到2500，达到QPS大于1000的目标（不选get接口因为SeesionID无法随机获取）。
4. 2000并发固定用户(get接口)QPS达到10000，达到QPS大于1500的目标。
5. 2000并发随机用户(login接口)QPS达到2500，达到QPS大于800的目标。


## 压测数据

|形式_接口_并发量|QPS|失败率|平均耗时|TP90|TP99|TP999|最长耗时|
|-----------|------|-----|-----|----|----|----|----|
|fix_get_200|11000|0|18ms|22ms|24ms|28ms|89ms|
|fix_get_2000|10000|0|190ms|235ms|253ms|596ms|2182ms|
|fix_login_200|2530|0|79ms|115ms|128ms|161ms|300ms|
|fix_login_2000|2400|0|800ms|1066ms|1401ms|2775ms|3063ms|
|rand_login_200|2530|0|79ms|111ms|124ms|152ms|377ms|
|rand_login_2000|2620|0|740ms|970ms|1356ms|2948ms|3051ms|

注：
压测过程中的tips:
1. 客户端发起的请求使用HTTP 1.1协议，使用长连接，不要使用短连接。
2. 客户端请求连接数可适当调大，避免出现`read: connection reset by peer`
3. 压测过程中关掉不必要的后台程序，避免产生 `Socket/File : too many open files(打开的文件过多) `、`EOF`报错。
4. 在Mac进行性能测试可能存在系统资源限制，可尝试下面指令：
```
sudo sysctl -w kern.ipc.somaxconn=2048
sudo sysctl -w kern.maxfiles=12288
ulimit -n 10000
```