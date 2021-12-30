package main

import (
	"fmt"
	"golang/shangguigu/customerManage/model"
	"golang/shangguigu/customerManage/service"
	"strconv"
)

type customerView struct {
	//定义必要字段
	key  string //接收用户输入
	loop bool   //表示是否循环的显示主菜单
	//增加一个字段customerService
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list() {
	//首先 获取到当前所有的客户信息(在切片中)
	customers := this.customerService.List()
	//显示
	fmt.Println("-------------------客户列表-------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].Getinfo())
	}
	fmt.Println("-----------------客户列表完成------------------")
}

//得到用户的输入信息，构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("-------------------添加客户---------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	//注意id号没有让用户直接输入，因为id是唯一的，需要系统分配
	customer := model.Newcustomer2(name, gender, age, phone, email)
	//调用
	if this.customerService.Add(*customer) {
		fmt.Println("-------------------添加完成---------------------")
	} else {
		fmt.Println("-------------------添加失败---------------------")

	}
}

//得到用户的输入，修改对应Id的客户
func (this *customerView) update() {
	fmt.Println("----------修改客户----------")
	fmt.Print("请选择修改客户的编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	customers := this.customerService.List()

	fmt.Printf("姓名(%v)：\n", customers[id-1].GetOneInfo("Name"))
	name := customers[id-1].GetOneInfo("Name")
	fmt.Scanln(&name)
	fmt.Printf("性别(%v)：\n", customers[id-1].GetOneInfo("Gender"))
	gender := customers[id-1].GetOneInfo("Gender")
	fmt.Scanln(&gender)
	fmt.Printf("年龄(%v)：\n", customers[id-1].GetOneInfo("Age"))
	age := customers[id-1].GetOneInfo("Age")
	fmt.Scanln(&age)
	fmt.Printf("电话(%v)：\n", customers[id-1].GetOneInfo("Phone"))
	phone := customers[id-1].GetOneInfo("Phone")
	fmt.Scanln(&phone)
	fmt.Printf("邮箱(%v)：\n", customers[id-1].GetOneInfo("Email"))
	email := customers[id-1].GetOneInfo("Email")
	fmt.Scanln(&email)

	customer := model.Newcustomer(id, name, gender, strconv.Atoi(age), phone, email)

	if this.customerService.Update(*customer) {
		fmt.Println("----------修改成功----------")
	} else {
		fmt.Println("----------修改失败----------")
	}
}

//得到用户的输入id，删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("-------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N)：")
labels:
	for {
		choice := ""
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" {
			//调用customerService 的Delete方法
			if this.customerService.Delete(id) {
				fmt.Println("-------------------删除完成---------------------")
				break labels
			} else {
				fmt.Println("删除失败，输入的id号不存在,请重新输入id号")
				break labels
			}
		} else if choice == "n" || choice == "N" {
			fmt.Println("不进行删除")
			break labels
		} else {
			fmt.Println("请输入Y/N,y/n")
		}
	}
}

//退出软件
func (this *customerView) exit() {
	fmt.Println("确认是否退出（Y/N）：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}

		fmt.Println("你的输入有误，确认是否退出（Y/N）：")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}

}

//显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("-------------------客户信息管理软件-------------------")
		fmt.Println("-------------------1 添 加 客 户---------------------")
		fmt.Println("-------------------2 修 改 客 户---------------------")
		fmt.Println("-------------------3 删 除 客 户---------------------")
		fmt.Println("-------------------4 客 户 列 表---------------------")
		fmt.Println("-------------------5  退   出  ----------------------")
		fmt.Println("请选择(1-5)")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("你的输入有误，请输入（1-5）！！！")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统。。。")
}

func main() {
	//在主函数中，创建一个customerView，并运行显示主菜单。。。
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//这里完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}
