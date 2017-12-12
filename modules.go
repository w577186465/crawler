package crawler

import (
	"fmt"
)

// 创建模块 参数：name模块名称，model信息库Model
func CreateModule(name string, model interface{}) {
	dataTable := fmt.Sprintf("%s_data", name)  // 信息库表名
	linkTable := fmt.Sprintf("%s_links", name) // 链接库表名

	// 创建信息库
	if err := DB.Table(dataTable).CreateTable(model).Error; err != nil {
		fmt.Println(err)
	}

	CreateLinkTable(linkTable) // 创建链接库

	fmt.Printf("%s模块创建成功\r\n", name)
}

// 创建链接库
func CreateLinkTable(name string) {
	DB.Table(name).CreateTable(&Link{})
}
