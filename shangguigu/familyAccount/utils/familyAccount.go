package utils

import "fmt"

type familyAccount struct {
	//声明字段
	key     string
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义个变量，记录是否有收支的行为
	flag bool
	//收支的详情使用字符串来记录
	//当有收支时，只需要对details 进行拼接处理即可
	details string
	loop    bool
	user    string
	pwd     string
}

//编写一个工厂模式的构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户余额\t收支金额\t说明",
		user:    "",
		pwd:     "",
	}
}

//将显示明细写成一个方法
func (this *familyAccount) showDetails() {
	fmt.Println("---------------当前收支明细记录---------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细。。。来一笔吧！")
	}
}

//将登记收入写成一个方法
func (this *familyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//将登记支出写成一个方法
func (this *familyAccount) pay() {
	fmt.Println("登记支出")
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	//这里需要做一个必要的判断
	if this.money > this.balance {
		fmt.Println("余额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//将退出系统也写成一个方法
func (this *familyAccount) exit() {
	fmt.Println("你确定要退出吗？y/n")
	var choice string
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入")
	}
	if choice == "y" {
		this.loop = false
	}
}

//将转账功能写成一个方法
func (this *familyAccount) transfer() {
	fmt.Println("登记转账")
	fmt.Println("本次转账金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//添加一个验证方法
func (this *familyAccount) auth() {
	for {
		fmt.Println("请输入你的用户名：")
		fmt.Scanln(&this.user)
		fmt.Println("请输入你的密码：")
		fmt.Scanln(&this.pwd)
		if this.user == "root" && this.pwd == "123456" {
			fmt.Println("登陆成功！")
			return
		} else {
			fmt.Println("用户名或者密码错误！登陆失败！")
		}

	}
}

//给该结构体绑定相应的方法
func (this *familyAccount) MainMenu() {
	this.auth()
	for {
		fmt.Println("\n---------------家庭收支记账软件---------------")
		fmt.Println("---------------1 收支明细---------------")
		fmt.Println("---------------2 登记收入---------------")
		fmt.Println("---------------3 登记支出---------------")
		fmt.Println("---------------4 退出软件---------------")
		fmt.Println("---------------5 登记转账---------------")
		fmt.Println("请选择(1-5)：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		case "5":
			this.transfer()
		default:
			fmt.Println("输入错误！请输入1-5的选项！")
		}
		if !this.loop {
			break
		}

	}

}
