package model

// mac平台下交叉编译linux可执行文件，不能使用 "gorm.io/driver/sqlite"，会报错 ref：https://blog.xiaoz.org/archives/18195
import (
	"dtext-go/utils"
	"encoding/json"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

var DB *gorm.DB

type Page struct {
	gorm.Model
	PageName    string // 页面名称
	PageContent string // 页面内容
	ReadOnly    bool   `gorm:"default:false"` // 仅读
}

func init() {
	db, err := gorm.Open(sqlite.Open(utils.RelativePath("dtext.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Page{})
	DB = db
}

func WriteDefaultPages(defaultPagesFilePath string) {
	// 从json文件中读取json内容
	var file, err = os.Open(defaultPagesFilePath)
	if err != nil {
		return
	}
	// 释放文件对象
	defer file.Close()

	jsonFile, _ := ioutil.ReadAll(file)
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonFile), &result)
	for key, value := range result {
		page := Page{PageName: key, PageContent: value.(string), ReadOnly: true}
		var tmpPage Page
		DB.Where("page_name = ?", key).Find(&tmpPage)
		if tmpPage.ID != 0 {
			DB.Where("page_name = ?", key).Updates(&page)
		} else {
			DB.Create(&page)
		}
	}
}

func DeletePage(pageName string) error {
	// 删除 pageName对应的page
	rst := DB.Where("page_name = ?", pageName).Delete(&Page{})
	return rst.Error
}

func GetPage(pageName string) Page {
	// 获取 pageName 对应的page
	if pageName == "" {
		pageName = "index"
	}
	var page Page
	rst := DB.Where("page_name = ?", pageName).Find(&page)
	if rst.Error != nil {
		panic(fmt.Sprintf("尝试从数据库中获取 %s 失败", pageName))
	}
	return page
}

func WritePage(pageName string, pageContent string) bool {
	original := GetPage(pageName)
	// 新页面
	if original.ID == 0 {
		page := Page{
			PageName:    pageName,
			PageContent: pageContent,
		}
		rst := DB.Create(&page)
		if rst.Error != nil {
			panic(fmt.Sprintf("新建 %s 时出错; 错误信息: %s", pageName, rst.Error))
		}
		return true
	}

	// 仅读页面
	if original.ReadOnly {
		return false
	}

	// 更新页面
	original.PageContent = pageContent
	rst := DB.Updates(&original)
	if rst.Error != nil {
		panic(fmt.Sprintf("更新 %s 时出错; 错误信息: %s", pageName, rst.Error))
	}
	return true
}
