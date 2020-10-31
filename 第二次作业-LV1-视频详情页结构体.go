package main

import "fmt"

type author struct {
	Name string				//名字
	VIP bool 				//是否是高贵的大会员
	Icon string 			//头像
	Signature string 		//签名
	Focus int 				//关注人数
}

func main()  {
	author1:=author{
		 Name: "老番茄",
		  VIP:true,
		 Icon: "一个红色的番茄",
	Signature:"新浪微博:_老番茄_",
	    Focus:1324300,

	}
fmt.Print("Name:",author1.Name )

fmt.Print("VIP:",author1.VIP)

fmt.Print("Icon:",author1.Icon)

fmt.Print("Signature:",author1.Signature)

fmt.Print("focus:",author1.Focus)

}