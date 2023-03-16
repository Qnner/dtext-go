[English Version](https://github.com/Qnner/dtext-go/blob/main/README-en.md)

# dtext
dtext 是一款在线文本同步软件，dtext-go是由go语言编写的版本；您**无需借助微信、QQ等聊天软件**，使用任何可以上网的设备，即可同步文本信息。

dtext会架设一个http网站，当网站运行起来后，任何可以访问这个网站的人，仅需访问/pageName/w 并编辑网页的内容,
都可以对几乎任何页面做修改(仅读页面除外，默认包括index 和 help页面)并生效。
当您以/pageName的形式访问的时候，默认为仅读模式，防止您误操作页面的内容。

## 使用指引
- demo： [https://go.dtext.cn](https://go.dtext.cn)
- 编译部署：
```shell
# 获取项目源码
git clone https://github.com/Qnner/dtext-go.git

# 编译
go build -o dtext

# 运行
./dtext
```



## 工作原理
当您以 写 模式打开页面时，您实际上是打开了一个铺满全屏幕的输入框，js脚本会每秒钟读取文本框内容一次，并向服务器同步您的更改。

## 项目依赖
- [gin](https://github.com/gin-gonic) 号称Golang中最快的web框架
- [gorm](https://github.com/go-gorm/gorm) Golang中优秀的ORM框架

## 许可证
本项目为作者学习Golang创作，项目遵从 [MIT](https://github.com/Qnner/dtext-go/blob/master/LICENSE) 协议
