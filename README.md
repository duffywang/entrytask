# entrytask
first entry task

## cmd
支持编译不同二进制程序的包，比如Restful路由程序，需要相关router, handler包和main入口包。

## internal
项目内部使用的包，包括crud, service(facade)和业务逻辑的包。

## pkg
如果你把代码包放在根目录的pkg下，其他项目是可以直接导入pkg下的代码包的，即这里的代码包是开放的，当然你的项目本身也可以直接访问的。
如果你的项目是一个开源的并且让其他人使用你封装的一些函数等，放在pkg中是合适的。





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
*golint
goreturns
dic: 扩展调试
gopls
gocode-gomod 
