# entrytask
first entry task

## cmd
支持编译不同二进制程序的包，比如Restful路由程序，需要相关router, handler包和main入口包。

## internal
项目内部使用的包，包括crud, service(facade)和业务逻辑的包。


## pkg
如果你把代码包放在根目录的pkg下，其他项目是可以直接导入pkg下的代码包的，即这里的代码包是开放的，当然你的项目本身也可以直接访问的。
如果你的项目是一个开源的并且让其他人使用你封装的一些函数等，放在pkg中是合适的。