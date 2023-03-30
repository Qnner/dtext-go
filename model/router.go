package model

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

var Server *gin.Engine

type msg struct {
	Msg string `form:"msg"`
}

func InitServer(temp embed.FS, static embed.FS) *gin.Engine {

	// release mode
	// gin.SetMode(gin.ReleaseMode)


	Server = gin.Default()

	Server.Any("/static/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(static))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})
	templ := template.Must(template.New("").ParseFS(temp, "templates/*.html"))
	Server.SetHTMLTemplate(templ)

	/* 使用非 embed 形式的文件管理系统
	Server.LoadHTMLGlob("templates/*")
	Server.StaticFS("/static", http.Dir("./static"))
	*/

	Server.GET("/", GetPageController) // 根路径 返回index的内容
	Server.GET("/:pageName", GetPageController)
	Server.GET("/:pageName/:option", GetPageWithOptionController)
	Server.POST("/:pageName/:option", PostPageWithOptionController)
	return Server
}

func GetPageController(ctx *gin.Context) {
	pageName := ctx.Param("pageName")
	pageContent := GetPage(pageName)
	ctx.HTML(http.StatusOK, "template.read.html", gin.H{
		"page_title":   pageName,
		"page_content": pageContent.PageContent,
	})

}

func GetPageWithOptionController(ctx *gin.Context) {
	pageName := ctx.Param("pageName")
	option := ctx.Param("option")
	pageContent := GetPage(pageName)
	if pageContent.ReadOnly {
		GetPageController(ctx)
		return
	}
	switch option {
	case "w":
		ctx.HTML(http.StatusOK, "template.write.html", gin.H{
			"page_title":   pageName,
			"page_content": pageContent.PageContent,
		})
		return
	case "write":
		ctx.HTML(http.StatusOK, "template.write.html", gin.H{
			"page_title":   pageName,
			"page_content": pageContent.PageContent,
		})
		return
	default:
		GetPageController(ctx)
		return
	}

}

func PostPageWithOptionController(ctx *gin.Context) {
	pageName := ctx.Param("pageName")
	option := ctx.Param("option")
	switch option {
	default:
		ctx.String(http.StatusOK, fmt.Sprintf("您正以POST的方式，以 %s 的方式访问 %s 页面，如果您想尝试写这个页面，"+
			"请使用 'w' 或者 'write' 模式", option, pageName))
		return
	case "w":
		postPageWriteOption(ctx)
		return
	case "write":
		postPageWriteOption(ctx)
		return
	}

}

func postPageWriteOption(ctx *gin.Context) {
	pageName := ctx.Param("pageName")
	var pc msg
	if ctx.ShouldBind(&pc) != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("参数错误！请检查您的参数，参数格式：%s", pc))
	} else {
		ret := WritePage(pageName, pc.Msg)
		if ret {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "success",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    -1,
				"message": "something error",
			})
		}

	}
}
