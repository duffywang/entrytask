# 一、entrytask项目简介
实现一个用户管理系统，用户可以登录、拉取和编辑他们的用户信息。用户可以在Web页面输入用户名和密码登录，后端系统负责校验用户身份，成功登录后展示用户的相关信息，可以修改用户昵称和用户头像。

# 二、系统设计
## 2.1 项目包结构
**benchmark:**
压测方法介绍，存放性能测试数据。

**cmd:**
包括http服务端、rpc服务端和客户端的配置读取，应用构建部署入口。
- http-server:HTTP服务端部署入口。
- rpc-server:RPC服务端部署入口。

**configs:**
项目全局配置数据，包括HTTP服务配置、连接数据库配置、Redis缓存配置以及gRPC服务配置。

**global:**
项目使用到的数据库客户端、redis客户端以及gRPC客户端。

**img：**
存放README.md文档中使用到的图片。

**internal：**
项目内部核心逻辑，包括前端展示层、Web API层，Service层、Dao数据持久化层，包括具体业务逻辑。
- constant:项目中用到的常量统一放在这里。
- web:定义路由匹配逻辑，并发起对应的RPC调用。
- service:组合各种数据访问构建的业务逻辑。
- dao:数据读写层，数据库和缓存全部放在这层统一处理。
- models:放对应的“存储层”的结构体，与存储中字段一一映射。

**log：**
记录应用日志文件
- http-server:记录HTTP服务端日志。
- rpc-server:记录RPC服务端日志。

**pkg：**
项目中使用到的中间件、工具。
- middleware:SessionID校验、登录校验中间件、。
- response:对返回结果进行封装。
- setting:定义以及读取配置工具。
- utils:包括文件处理工具、哈希加密工具。

**proto：**
gRPC服务使用proto buffer序列化方式，包中包括定义的用户信息pb文件和文件新pb文件，以及使用protoc指令生成的文件。

**upload：**
保存用户上传的头像图片。

**view：**
用户系统前端HTML页面代码。


## 2.2系统架构设计
![系统架构设计](https://github.com/duffywang/entrytask/blob/main/img/硬件架构.png#pic_center)

主要由5部分构成：
1. Web客户端浏览器是请求发起端，用户可以在前端页面上进行登录、注册操作，发送相应的HTTP协议GET、POST请求至Web服务器。
2. Web服务端接收Web客户端发起的请求，根据路由规则匹配到对应的处理逻辑上，Web服务端会发起RPC请求，HTTP Server使用gin框架。
3. RPC服务端注册提供具体业务处理逻辑，查询数据库、缓存中的数据，RPC服务端使用gRPC框架搭建，序列化方式使用proto buf。
4. RPC服务端需要存储和读取用户登录的SessionID，使用Redis缓存。
5. 数据库存储着用户信息表，提供查询用户信息、新增用户、更新用户信息等数据操作，使用MySQL数据库。


## 2.3 业务架构设计
![业务架构设计](https://github.com/duffywang/entrytask/blob/main/img/业务架构.png#pic_center)

## 2.4 接口设计
系统一共对外提供了5个接口，分别是：
1. 用户登录接口
2. 用户注册接口
3. 获取用户接口
4. 编辑用户接口
5. 上传图片接口

### 登录接口 api/user/login POST

**输入参数**
|    字段名   |类型             |是否必填|含义 |
|----------------|----------------|-----------------|--|
|username|string           |是         |用户名|
|password    |string         |是         |密码|

**返回参数**
|   通用字段 | 业务字段  |类型      |是否可为空|含义 |
|---------|-------|----------------|-----------------|--|
|code | |int           |否         |返回码|
| msg| |string           |否         |返回信息|
|data | |any          |是         | 具体业务数据|
| |username|string           |否         |用户名|
| |nickname    |string         |是        |昵称|
| |profile_pic    |string     |是         |用户头像|
| |session_id    |string         |是         |sessionID|

**请求示例**
```
curl -H "Content-Type:application/json" -X POST -d '{"username":"test4","password":"1234567"}' 'http://127.0.0.1:8080/api/user/login'
```

**返回示例**
```
//success
{
    "code":0,
    "data":{
        "username":"test4",
        "nickname":"nicntest4",
        "profile_pic":"",
        "session_id":"0465888d-bf04-42ca-82ad-0156facffba4"
    },
    "msg":"success"
}


//fail
{
    "code":300000,
    "data":"密码错误，请重试",
    "msg":"User Login Error"
}
```


### 用户注册接口 api/user/register POST

**输入参数**
|    字段名   |类型             |是否必填|含义 |
|----------------|----------------|-----------------|--|
|username|string           |是         |用户名|
|password  |string         |是         |密码|
|nickname  |string         |否       |昵称|
|profile_pic |string         |否         |用户头像|

**返回参数**
|   通用字段   |类型      |是否可为空|含义 |
|---------|----------------------|-----------------|--|
|code |int           |否         |返回码|
| msg |string           |否         |返回信息|
|data  |any          |是        | 具体业务数据|

**请求示例**
```
curl -H "Content-Type:application/json" -X POST -d '{"username":"test4","password":"1234567","nickname":"nicntest4","profile_pic":"xixi"}' 'http://127.0.0.1:8080/api/user/register'
```

**返回示例**
```
 //success
{
    "code":0,
    "data":{},
    "msg":"success"
}
//fail
{
    "code":300001,
    "data":"用户名已存在",
    "msg":"User Register Error"
}
```

### 获取用户信息接口 api/user/get GET

**输入参数**
|    字段名   |类型             |是否必填|含义 |
|----------------|----------------|-----------------|--|
|session_id|string           |是         |sessionID|

**返回参数**
|   通用字段 | 业务字段  |类型      |是否可为空|含义 |
|---------|-------|----------------|-----------------|--|
|code | |int           |否         |返回码|
| msg| |string         |否       |返回信息|
|data | |any          |是         | 具体业务数据|
| |username|string           |否         |用户名|
| |nickname|string         |是        |昵称|
| |profile_pic |string         |是         |用户头像|

**请求示例**
```
curl --location --request GET --cookie 'session_id=86261c69-61c1-42d2-bc69-a28610e93a9b' 'http://127.0.0.1:8080/api/user/get'
```

**返回示例**
```
//success
{
    "code":0,
    "data":{
        "username":"test",
        "nickname":"testnick",
        "profile_pic":" "
    },
    "msg":"success"
}

//fail
{
    "code":300002,
    "data":"SessionID错误",
    "msg":"User Get Error"
}
```

### 编辑用户信息接口 api/user/edit POST

**输入参数**
|    字段名   |类型             |是否必填|含义 |
|----------------|----------------|-----------------|--|
|nickname  |string         |否       |昵称|
|profile_pic|string         |否    |用户头像|
|session_id|string           |是      |sessionID|

**返回参数**
|   通用字段   |类型      |是否可为空|含义 |
|---------|----------------------|-----------------|--|
|code |int           |否         |返回码|

**请求示例**
```
curl -H "Content-Type:application/json" -X POST -d '{"nickname":"testedit","profilepic":"hahaedit"}' 'http://localhost:8080/api/user/edit' --cookie 'session_id=ffb1fbdd-1784-438c-a8ef-af5cbc1a5022'
```

**返回示例**
```
//success
{
    "code":0,
    "data":{
    },
    "msg":"success"
}

//fail
{
    "code":300003,
    "msg":"User Edit Error"
}
```

### 上传图片接口 api/file/update POST
**输入参数**
|    字段名   |类型       |是否必填|含义|
|------------|----------|-------|--|
|file        |file      |是     |文件|


**返回参数**
|   通用字段 | 业务字段  |类型      |是否可为空|含义|
|---------|-------|--------------|----------|--|
|code |    |int           |否         |返回码|
| msg |    |string         |否        |返回信息|
|data |    |any           |是         |具体业务数据|
|     |filername|string    |是       |文件名|
|     |fileurl  |string   |是        |文件地址|

请求在页面上直接上传文件

**返回示例**
```
//success
{
    "code":0,
    "data":{
        "filename":"51ebd8c445683e0520a5f79e8b999c1d.png",
        "fileurl":"http://localhost:8080/static/51ebd8c445683e0520a5f79e8b999c1d.png"
    },
    "msg":"success"
}
//fail
{
    "code":300004,
    "msg":"File Upload Error"
}
```


# 四、项目部署
部署涉及到HTTP服务器和RPC服务器、数据库MySQL服务器和缓存Rdis，服务部署均在个人PC上，没有使用虚拟机。下面分别讲解部署方法：
## http服务器
启动HTTP服务器 
```
go run ./cmd/http-server/main.go
```

## rpc服务器
启动rpc服务器
```
go run ./cmd/rpc-server/main.go
```

## 数据库MySQL
官网下载MySQL，进入MySQL下载目录目录，MySQL服务器启动指令
```
sudo mysql.server start   //启动MySQL服务器
```
登录MySQL服务器，需要输入数据库密码
```
mysql -uroot -p
```
创建数据库entrytask
```
CREATE DATABASE `entrytask`
```
创建用户信息表user_table
```
CREATE TABLE `user_table` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) COLLATE UTF8MB4_UNICODE_CI NOT NULL,
    `nickname` VARCHAR(64) COLLATE UTF8MB4_UNICODE_CI NOT NULL,
    `profile_pic` VARCHAR(1024) COLLATE UTF8MB4_UNICODE_CI NOT NULL,
    `password` VARCHAR(64) CHARACTER SET UTF8MB4 COLLATE UTF8MB4_UNICODE_CI NOT NULL,
    `status` TINYINT UNSIGNED DEFAULT 0 NOT NULL,
    `create_time` INT UNSIGNED NOT NULL,
    `update_time` INT UNSIGNED NOT NULL,
    PRIMARY KEY(`id`),
    UNIQUE KEY `uniq_username` (`username`) 
    )ENGINE=INNODB DEFAULT CHARSET=UTF8MB4 COLLATE=UTF8MB4_UNICODE_CI;
```
user_table插入一条用户信息数据
```
mysql> insert into user_table
    -> (username,password,nickname,profile_pic,status,create_time,update_time)
    -> values
    -> ('test1','123456','testnick','haha',0,1651308516,1651308516);
```

表中插入10,000,000条用户账号信息
首先开启创建存储函数
```
SET GLOBAL log_bin_trust_function_creators=TRUE;
```
数据批量插入存储函数，函数名为`insert_user_table_v2`
```
DELIMITER '$';
CREATE FUNCTION insert_user_table_v2()
RETURNS INT
BEGIN
	DECLARE num INT DEFAULT 10000000;
	DECLARE i INT DEFAULT 1;
	WHILE i < num DO
	INSERT INTO user_table(username,nickname,profile_pic,password,status,create_time,update_time)VALUES (CONCAT('test',i),'auto_nickname','  ','9c4bd805568b48f15bb0618fe5ba4461',0,1651308518,1651308518);
	SET i = i + 1;
	END WHILE;
	RETURN i;
END;
$
```
触发函数执行，可能需要几分钟
```
SELECT insert_user_table_v2(); $ //别忘记结束符$
```
数据库中成功插入了10,000,000条数据
```
mysql>select count(*) from user_table;
+----------+
| count(*) |
+----------+
| 10000007 |
+----------+
```

## Redis
官网下载redis，进入redis所在目录，Mac系统在usr/local/opt/redis/bin，执行下面命令启动redis服务端。
```
redis-server
```
同理，在另一个终端页面进入usr/local/opt/redis/bin目录下，执行下面命令启动redis客户端。
```
redis-cli
```

# QA
Q1:如何做错误处理？
针对依赖核心组件，启动项目会先初始化数据库、缓存组件，如果启动时发生异常，执行`panic(err)`中断。
```
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
        //values
	)))
	if err != nil {
		panic(err)
	}
```

针对业务产生错误，对错误分类，在上面接口设计中看到登录异常和注册异常返回的错误信息不同，帮助开发人员更高效定位问题
|    code   |类型       |含义|
|------------|----------|--|
|200000      |InvalidParamsError |入参异常|
|300000      |UserLoginError |密码错误|
|300001      |UserRegisterError |用户名已存在|
|300002      |UserGetError |获取不到用户|
|300003      |UserEditError |编辑信息异常|
|300004      |FileUploadError |文件上传异常|

针对一些特殊场景，返回err不为空不一定是异常，比如注册用户信息，首先会查询注册的用户名在表中是否已存在，如果表中查不到该用户名，会返回`ErrRecordNotFound`，说明用户可以注册，否则无法注册
```
	_, err := svc.dao.GetUserInfo(request.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		//说明已存在
		return nil, errors.New("RegisterUser Fail : Username Exist")
	}
```

Q2:缓存一致性如何保证？


Q3:鉴权机制如何做的？
1. 用户在浏览器端填写用户名和密码登录
2. 服务端对用户名和密码校验通过后会生成一份保存当前用户相关信息的session数据和一个与之对应的的标示（通常称为session_id）,session_id生成方式选择```sessionID := uuid.NewV4()```，系统以username-session_id的key-value形式存储于redis中，设定一定有效期（30分钟）。
3. 用户在登录成功后，服务端返回响应时将session_id写入用户浏览器的Cookie。
4. 后续用户来自该浏览器的每次请求都后自动携带包含session_id的Cookie。
5. 服务端通过请求中的session_id就能找到之前保存的该用户的那份session数据，从而获取该用户的相关信息。

Q4:SQL注入/*
>> SQL注入是一种注入攻击手段，通过执行恶意SQL语句，将任意SQL代码插入到数据库查询，从而使攻击者完全控制Web应用程序后台的数据库服务器。避免SQL注入的一般原则是，不信任用户提交的数据。
我们参数合法性判断和参数化查询的方法避免MySQL注入。
采用前置参数合法性判断，举个例子如果username为`1=1;drop table users;`
`“SELECT * FROM users WHERE 1=1;drop table users;”`会导致MySQL注入发送，因此参数中不允许包含";"字段，如果含有直接返回参数异常。
```
	if i := strings.Index(param.Username,";");i != -1 {
		resp.ResponseError(constant.InvalidParamsError)
		return
	}
```
采用参数化查询的方法，先将SQL语句中可能被客户端控制的参数集进行编译，生成对应的临时变量集，再使用对应的设置方法，为临时变量集里面的元素进行赋值。
```
    db = db.Where("username = ?", u.Username)
```

Q5:密码等敏感数据处理？
首先了解下常见的敏感数据数据加密方法

|加密算法|加密方式|破解难度|
|--|-----------------|--|
|无|明文保存|易|
|AES|对称加密|  易|
|MD5、SHA1|单向Hash |中 |
|salt+MD5|双向Hash  |中 |
|bcrypt| 多次Hash | 难|

go语言中`golang.org/x/crypto/bcrypt`包中提供了bcrypt加密方法，但是据性能测试bcrypt算法对性能影响很大，200并发请求情况下qps低于100。
```
    //注册时，数据库存入加密后的密码
	pwd, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

    //登录时，校验数据库中密码与输入密码是否相同
	err = bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(r.Password))

```
项目中为了在性能和安全性中找到平衡，最终选择使用salt+用户密码进行MD5哈希方式，取哈希结果存储在数据库中。
```
func Hash(password string)string {
	hash := md5.Sum([]byte("salt"+password))
	return hex.EncodeToString(hash[:])
}
```

# 参考
1. https://github.com/gin-gonic/gin
2. https://github.com/grpc/grpc-go
3. https://github.com/protocolbuffers/protobuf
4. https://github.com/go-gorm/gorm
5. https://github.com/link1st/go-stress-testing
6. https://github.com/go-redis/redis/v8

