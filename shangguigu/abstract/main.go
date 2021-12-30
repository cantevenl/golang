package main

import "fmt"

//定义一个结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//方法
//1.存款
func (account *Account) Deposite(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码错误")
		return
	}

	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance += money
	fmt.Println("存款成功，目前余额\n", account.Balance)
}

//取款
func (account *Account) WithDraw(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码错误")
		return
	}

	//看看取款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功，目前余额\n", account.Balance)
}

//查询余额
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码错误")
		return
	}

	fmt.Printf("你的账号是%v,你的余额为%v\n\n", account.AccountNo, account.Balance)
}
func main() {
	account := Account{
		"zhuzhu666",
		"xk123",
		100,
	}
	for {
		var opertor string
		var pwd string
		var money float64
		fmt.Println("欢迎来到银行：\n按1进行存款\n按2进行取款\n按3进行查询\n按n退出银行")
		fmt.Scanln(&opertor)
		switch opertor {
		case "1":
			fmt.Println("请输入密码")
			fmt.Scanln(&pwd)
			fmt.Println("请输入需要存入多少钱")
			fmt.Scanln(&money)
			account.Deposite(money, pwd)
		case "2":
			fmt.Println("请输入密码")
			fmt.Scanln(&pwd)
			fmt.Println("请输入需要取多少钱")
			fmt.Scanln(&money)
			account.WithDraw(money, pwd)
		case "3":
			fmt.Println("请输入密码")
			fmt.Scanln(&pwd)
			account.Query(pwd)
		case "n":
			fmt.Println("退出银行")
			return
		default:
			fmt.Println("输入有误")

		}

	}
}
