# 一、entrytask项目简介
实现一个用户管理系统，用户可以登录、拉取和编辑他们的用户信息。用户可以在Web页面输入用户名和密码登录，后端系统负责校验用户身份，成功登录后展示用户的相关信息，可以修改用户昵称和用户头像，。

TODO：参考标准的README.md文档规范是怎么写的

# 二、项目包结构
## cmd
支持编译不同二进制程序的包，需要相关router, handler包和main入口包。
包括http服务端、rpc服务端和客户端的配置读取，构建部署

## internal
项目内部核心逻辑，包括前端展示层、Web API层，Service层、Dao数据持久化层，包括具体业务逻辑

## pkg
如果你把代码包放在根目录的pkg下，其他项目是可以直接导入pkg下的代码包的，即这里的代码包是开放的，当然你的项目本身也可以直接访问的。
如果你的项目是一个开源的并且让其他人使用你封装的一些函数等，放在pkg中是合适的。

## benchmark
性能测试数据，要求如下
1. 数据库中必须有10000000条用户信息
2. 返回结果是正确的
3. 每个请求都要包含RPC调用和MySQL或Redis访问
4. 200并发固定用户，HTTP API QPS大于3000；200并发随机用户HTTP API QPS大于1000
5. 2000并发固定用户，HTTP API QPS大于1500；2000并发随机用户HTTP API QPS大于800

## configs
项目全局配置数据，包括HTTP服务配置、连接数据库配置、Redis缓存配置以及gRPC服务配置

## global
项目使用到的数据库客户端、redis客户端以及gRPC客户端

## proto
gRPC服务使用proto buffer序列化方式，包中包括定义的用户信息pb文件和文件新pb文件，以及使用protoc指令生成的文件，举个例子
```
protoc --go-grpc_out=. user.proto
protoc --go_out=. user.proto
```
具体可参考官方文档： https://github.com/protocolbuffers/protobuf

## upload
保存用户上传的头像图片

## log
记录日志文件，分为HTTP服务端日志和RPC服务端日志


# 三、部署
## 数据库MySQL
启动相关指令
```
sudo mysql.server start   //启动MySQL服务器
sudo mysql.server stop    //关闭MySQL服务器
sudo mysql.server restart //重启MySQL服务器
```

登录，需要输入数据库密码
```
mysql -uroot -p
```
常用指令
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
## Redis
进入redis文件usr/local/opt/redis/bin目录下，启动redis服务端
```
redis-server
```
进入redis文件usr/local/opt/redis/bin目录下，启动redis客户端
```
redis-cli
```
# 四、 


使用的标准库 fmt error flag log strov
使用的第三方库

使用的插件
Go
Shades of Purple
vscode-icons
vscode-protos
Clang-Format

gocode 代码自动完成
gopkgs 未导入的软件包提供自动补全功能
go-outline 文档大纲功能？
go-symbol ？
guru 查找参考和查找接口实现功能
gorename: 重命名功能
gotests: 为Go:Generate Unit Tests 指令提供支持
*impl: 为Go:Generate Interface Stubs 命令提供支持 
*gomodifytags: 为Go:Add Tags to Struct Fields 和Remove Tags From Struct Fields命令
*fillstruct: 为Go:Fill struct命令的支持
*goplay: 为Go: Run on Go Playgrouond 命令提供支持
*golint 代码规范？何为通过go lint
goreturns
dic: 扩展调试
gopls
gocode-gomod 

# 参考