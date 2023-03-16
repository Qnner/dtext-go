[中文版本](https://github.com/Qnner/dtext-go/blob/main/README.md)

# dtext
dtext is an online text synchronization software, and dtext-go is a version written in the Go language. You don't need to rely on chat software like WeChat or Telegram. Using any device that can access the internet, you can synchronize text information.

dtext will set up an HTTP website. After the website is up and running, anyone who can access this website can make modifications to almost any page (except for read-only pages, which by default include the index and help pages) and make them effective by simply accessing /pageName/w and editing the webpage's content. When you access it in the form of /pageName, it is in read-only mode by default to prevent you from accidentally modifying the content of the page.

## How it works
When you open a page in "write" mode, you are actually opening a full-screen input box. The JavaScript script will read the contents of the text box once per second and synchronize your changes with the server.

## Usage
- demo： [https://go.dtext.cn](https://go.dtext.cn)
- Compile and deploy：
```shell
# Get project source code
git clone https://github.com/Qnner/dtext-go.git

# Compile
go build -o dtext

# deploy
./dtext
```

## Dependencies
- [gin](https://github.com/gin-gonic) Claimed to be the fastest web framework in Golang.
- [gorm](https://github.com/go-gorm/gorm) Excellent ORM frameworks in Golang

## LICENSE
This project was created by the author to learn Golang and follows the [MIT](https://github.com/Qnner/dtext-go/blob/main/LICENSE) license.

