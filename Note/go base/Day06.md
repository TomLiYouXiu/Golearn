# 依赖管理

在本阶段主要是使用别人的库

<img src="https://pic.imgdb.cn/item/65691ef6c458853aef5b2a7a.jpg" style="zoom:50%;" />

## GOPATH

<img src="https://pic.imgdb.cn/item/65691fd0c458853aef5c766d.jpg" style="zoom:50%;" />

## GO VENDOR

<img src="https://pic.imgdb.cn/item/65692806c458853aef6e297a.jpg" style="zoom:50%;" />

## GO MOD

go mod tidy 清除无用的依赖

go get -u xxx 下载指定的依赖

go get -u xxx@vXXX 指定下载版本

也可以直接import 会自动导入   

**项目迁移命令**

~~~
go mod init xxx
go build ./... //下载所有依赖
~~~

~~~
glide.yaml
~~~

<img src="https://pic.imgdb.cn/item/65692ff9c458853aef7eb5fb.jpg" style="zoom: 50%;" />

# 文件整理

