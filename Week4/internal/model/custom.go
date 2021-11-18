package model

import "fmt"

//生命一个customer结构，表示一个客户信息

type Customer struct {
	Id    int
	Name  string
	Phone string
	Email string
}

//编写一个工厂模式，返回一个customer的示例
func NewCustomer(id int, name string, phone string, email string) Customer {
	return Customer{
		Id:    id,
		Name:  name,
		Phone: phone,
		Email: email,
	}
}
func NewCustomer2(name string, phone string, email string) Customer {
	return Customer{
		Name:  name,
		Phone: phone,
		Email: email,
	}
}

//返回用户的信息，格式化的字符串
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Phone, this.Email)
	return info
}
