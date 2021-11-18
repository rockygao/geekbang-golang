package biz

import (
	"fmt"

	"Week4/internal/model"
	"Week4/internal/service"
)

type CustomerView struct {

	//定义必要字段
	Key  string //接收用户输入
	Loop bool   //表示是否循环显示主菜单
	//增加一个字段 customerService
	CustomerService *service.CustomerService
}

//根据用户输入的id  删除
func (this *CustomerView) Delete() {
	fmt.Println("----------删除客户----------")
	fmt.Println("请输入要删除的客户编号Id(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除
	}

	//确定要删除吗
	fmt.Println("确定要删除吗 确认输入 Y：")
	choice := ""
	fmt.Scanln(&choice)
	if choice != "y" && choice != "Y" {
		return
	}

	//调用删除
	if this.CustomerService.Delete(id) {
		fmt.Println("\n----------删除成功----------\n\n")
	} else {
		fmt.Println("\n----------id 不存在----------\n\n")
	}

}

//添加客户信息
func (this *CustomerView) Add() {
	//显示
	fmt.Println("----------添加客户----------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("手机号：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的customer实例
	customer := model.NewCustomer2(name, phone, email)

	//调用添加
	if this.CustomerService.Add(customer) {
		fmt.Println("\n----------添加成功----------\n\n")
	} else {
		fmt.Println("\n----------添加失败,已存在----------\n\n")
	}
}

//显示所有客户信息
func (this *CustomerView) List() {
	//首先，获取当前所有客户信息（在切片中）
	customers := this.CustomerService.List()
	//显示
	fmt.Println("----------客户列表----------")
	fmt.Println("Id\t姓名\t手机号\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("\n----------客户列表----------\n\n")
}

//退出
//根据Id 删除客户信息 从切片中删除
func (this *CustomerView) Exit() {
	fmt.Println("确定退出吗 （Y/N）:")
	for {
		fmt.Scanln(&this.Key)
		if this.Key == "Y" || this.Key == "y" || this.Key == "N" || this.Key == "n" {
			break
		}
		fmt.Println("输入有误，请输入（Y/N）")
	}
	if this.Key == "Y" || this.Key == "y" {
		this.Loop = false
	}
}

//显示主菜单
func (this *CustomerView) MainMenu() {
	for {
		fmt.Println("----------客户信息----------------")
		fmt.Println("----------1 添加----------------")
		fmt.Println("----------2 修改----------------")
		fmt.Println("----------3 删除----------------")
		fmt.Println("----------4 列表----------------")
		fmt.Println("----------5 退出----------------")
		fmt.Println()
		fmt.Print("请选择1-5:")

		fmt.Scanln(&this.Key)

		switch this.Key {
		case "1":
			this.Add()
		case "2":
			fmt.Println("修改")
		case "3":
			this.Delete()
		case "4":
			this.List()
		case "5":
			this.Exit()
		default:
			fmt.Print("输入有误，重新输入")
		}

		if !this.Loop {
			break
		}

	}
	fmt.Println("你退出了该系统")
}

// func main() {
// 	//在主函数中创建一个customerView，并运行显示主菜单
// 	customerView := CustomerView{
// 		Key:  "",
// 		Loop: true,
// 	}
// 	//这里完成对customerView结构体的customerService字段初始化
// 	customerView.CustomerService = service.NewCustomerService()
// 	//显示主菜单
// 	customerView.MainMenu()
// }
