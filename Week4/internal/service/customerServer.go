package service

import (
	"Week4/internal/model"
)

//该CustomerService ,完成对Customer的操作，包括
//增删改查
//会声明一个customer的切片
type CustomerService struct {
	Customers []model.Customer
	//声明一个字段，表示当前切片有多少个客户
	//该字段后面，还可以作为新客户的id+1
	CustomerNum int
}

//编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	//为了可以看到客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.CustomerNum = 1
	customer := model.NewCustomer(1, "张三", "13166888866", "test@163.com")
	customerService.Customers = append(customerService.Customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.Customers
}

//添加客户到customers切片
//！！！ 需要使用 引用 *CustomerService
func (this *CustomerService) Add(customer model.Customer) bool {
	index := this.FindCustomerByPhone(customer.Phone)
	if index == 1 {
		return false
	}
	//分配一个id的规则，就是添加的顺序
	this.CustomerNum++
	customer.Id = this.CustomerNum
	this.Customers = append(this.Customers, customer)
	return true
}

//根据Id 删除客户信息 从切片中删除
func (this *CustomerService) Delete(id int) bool {
	index := this.FindCustomerById(id)
	//如果等于-1 则不存在
	if index == -1 {
		return false
	}

	//如何同切片中删除  从0 到 index(不含) 增加 index+1 到最后
	this.Customers = append(this.Customers[:index], this.Customers[index+1:]...)
	return true

}

//根据Id 查找客户是否切片中，如果存在，则返回下表，不存在返回下表
func (this *CustomerService) FindCustomerById(id int) int {
	index := -1 //默认不存在
	//遍历this.customers 切片
	for i := 0; i < len(this.Customers); i++ {
		if this.Customers[i].Id == id {
			//找到了
			index = i
		}
	}
	return index

}

//根据手机号 查找客户是否切片中，如果存在，则返回下表，不存在返回下表
func (this *CustomerService) FindCustomerByPhone(phone string) int {
	index := -1 //默认不存在
	//遍历this.customers 切片
	for i := 0; i < len(this.Customers); i++ {
		if this.Customers[i].Phone == phone {
			//找到了
			index = i
		}
	}
	return index

}
