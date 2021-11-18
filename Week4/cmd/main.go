package main

import (
	"Week4/internal/biz"
	"Week4/internal/service"
)

func main() {
	//在主函数中创建一个customerView，并运行显示主菜单
	customerView := biz.CustomerView{
		Key:  "",
		Loop: true,
	}
	//这里完成对customerView结构体的customerService字段初始化
	customerView.CustomerService = service.NewCustomerService()
	//显示主菜单
	customerView.MainMenu()
}
