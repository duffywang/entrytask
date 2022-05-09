sh-3.2# ./go-stress-testing-mac -c 200 -n 40 -p curl/login.txt

 开始启动  并发数:200 请求数:40 请求参数:
request:
 form:http
 url:http://127.0.0.1:8080/api/user/login
 method:POST
 headers:map[Accept:text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9 Accept-Language:zh-CN,zh;q=0.9,en;q=0.8 Cache-Control:max-age=0 Connection:keep-alive Content-Type:application/x-www-form-urlencoded Cookie:_lxsdk_cuid=1761e65afb6c8-0a5332a8efff16-18346153-13c680-1761e65afb6c8; _lxsdk=1761e65afb6c8-0a5332a8efff16-18346153-13c680-1761e65afb6c8; Idea-d82114bc=95219ce2-c359-40a4-a167-5a9b4a296b11; moaDeviceId=1651CF34E26C52D29911A79D6FC97EFD; webNewUuid=8f427f110df783ffe6025357be63e783_1649322370760; JSESSIONID=DC6189C2C18E86727D7F5EDB10BF6710 Origin:http://localhost:8080 Referer:http://localhost:8080/api/index Sec-Fetch-Dest:document Sec-Fetch-Mode:navigate Sec-Fetch-Site:same-origin Sec-Fetch-User:?1 Upgrade-Insecure-Requests:1 User-Agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36 sec-ch-ua:" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100" sec-ch-ua-mobile:?0 sec-ch-ua-platform:"macOS"]
 data:username=test4&password=1234567
 verify:statusCode
 timeout:30s
 debug:false



─────┬───────┬───────┬───────┬────────┬────────┬────────┬────────┬────────┬────────┬────────
 耗时│ 并发数│ 成功数│ 失败数│   qps  │最长耗时│最短耗时│平均耗时│下载字节│字节每秒│ 错误码
─────┼───────┼───────┼───────┼────────┼────────┼────────┼────────┼────────┼────────┼────────
   1s│    200│   1138│      0│ 1576.35│  220.91│   45.72│  126.88│1,143,798│1,143,769│200:1138
   2s│    200│   2138│      0│ 1381.83│  248.57│   45.72│  144.74│2,148,798│1,074,410│200:2138
   3s│    200│   2938│      0│ 1253.75│  345.34│   45.72│  159.52│2,952,798│ 984,244│200:2938
   4s│    200│   3661│      0│ 1150.24│  345.34│   45.72│  173.88│3,679,413│ 919,837│200:3661
   5s│    200│   4261│      0│ 1057.91│  435.35│   45.72│  189.05│4,282,413│ 856,249│200:4261
   6s│    200│   4789│      0│  988.71│  473.07│   45.72│  202.28│4,813,053│ 802,197│200:4789
   7s│    200│   5254│      0│  929.44│  496.91│   45.72│  215.18│5,280,378│ 754,354│200:5254
   8s│    200│   5653│      0│  875.56│  569.35│   45.72│  228.43│5,681,373│ 710,124│200:5653
   9s│    200│   5980│      0│  828.78│  607.37│   45.72│  241.32│6,010,008│ 667,711│200:5980
  10s│    200│   6247│      0│  770.67│ 1023.71│   45.72│  259.51│6,278,343│ 627,831│200:6247
  11s│    200│   6715│      0│  749.67│ 1023.71│   45.72│  266.78│6,748,683│ 613,528│200:6715
  12s│    200│   7241│      0│  732.55│ 1023.71│   45.72│  273.02│7,277,313│ 606,386│200:7241
  13s│    200│   7641│      0│  720.36│ 1023.71│   45.72│  277.64│7,679,313│ 590,722│200:7641
  13s│    200│   8000│      0│  714.96│ 1023.71│    6.43│  279.74│8,040,108│ 604,430│200:8000


*************************  结果 stat  ****************************
处理协程数量: 200
请求总数（并发数*请求数 -c * -n）: 8000 总请求时间: 13.302 秒 successNum: 8000 failureNum: 0
*************************  结果 end   ****************************