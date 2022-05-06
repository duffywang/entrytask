# 一、entrytask项目简介
实现一个用户管理系统，用户可以登录、拉取和编辑他们的用户信息。用户可以在Web页面输入用户名和密码登录，后端系统负责校验用户身份，成功登录后展示用户的相关信息，可以修改用户昵称和用户头像。


# 二、项目包结构
## benchmark
存放性能测试数据。

## cmd
包括http服务端、rpc服务端和客户端的配置读取，构建部署入口。

## configs
项目全局配置数据，包括HTTP服务配置、连接数据库配置、Redis缓存配置以及gRPC服务配置。

## global
项目使用到的数据库客户端、redis客户端以及gRPC客户端。

## img
说明文档中使用到的图片。

## internal
项目内部核心逻辑，包括前端展示层、Web API层，Service层、Dao数据持久化层，包括具体业务逻辑。

## log
记录日志文件，分为HTTP服务端日志和RPC服务端日志。

## pkg
项目中使用到的中间件、工具。

## proto
gRPC服务使用proto buffer序列化方式，包中包括定义的用户信息pb文件和文件新pb文件，以及使用protoc指令生成的文件。

## upload
保存用户上传的头像图片。

## view
用户系统前端HTML页面代码。


# 三、功能



# 四、部署
## 数据库MySQL
MySQL服务器启动相关指令
```
sudo mysql.server start   //启动MySQL服务器
sudo mysql.server stop    //关闭MySQL服务器
sudo mysql.server restart //重启MySQL服务器
```

登录MySQL服务器，需要输入数据库密码
```
mysql -uroot -p
```
MySQL常用指令，帮助使用
```
show databases;       //列出所有的数据库
show tables;          //列出所有的数据表
use ${database_name}; //使用选定数据库
desc ${table_name};   //表结构说明
```
创建数据库entrytask和创建用户信息表user_table
```
CREATE DATABASE `entrytask`

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
进入redis文件usr/local/opt/redis/bin目录下，启动redis服务端
```
redis-server
```
进入redis文件usr/local/opt/redis/bin目录下，启动redis客户端
```
redis-cli
```

## http服务器
启动http服务器，可以设置flag参数-mode -port 
```
go run ./cmd/http-server/main.go
```

## rpc服务器
启动rpc服务器
```
go run ./cmd/rpc-server/main.go
```




# 参考
https://github.com/gin-gonic/gin
https://github.com/grpc/grpc-go
https://github.com/protocolbuffers/protobuf
https://github.com/go-gorm/gorm
https://github.com/link1st/go-stress-testing

