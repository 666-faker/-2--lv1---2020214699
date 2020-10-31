package main

import "fmt"

type person struct {
	Age int
	Name string
	Height int


}

func main()  {
	Person1 :=person{

			Age:17,
		    Name: "faker",
		    Height: 67,
	}
	Person2 :=person{
		Age:18,
		Name:"LBW",
		Height:89,
	}
	fmt.Println("世界第一中单",Person1)
	fmt.Println("五五开",Person2)
}
