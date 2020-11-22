package main//入口函数

import "fmt"

func main()  {
	var name string
	var password string
	var loginChance = 3
	for i:=1; i <= 3 ; i++{
		fmt.Println("请输入账号")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&password)
		if name == "faker" && password == "123456"{
			fmt.Println("密码正确，已登入")
			break
		}else{loginChance--
			fmt.Println("密码错误,请重新输入")
		}
		if loginChance == 0 {
			fmt.Println("你的账号已被冻结")
		}


	}


}


