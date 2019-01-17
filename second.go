package main

import "fmt"

//定义接口
type  Phone interface {
	call()
}


type  NokiaPthon struct {

}

func (nokiaphone NokiaPthon) call(){
	fmt.Println("I am Nokia,I can sall you ")

}

type Iphone struct {

}

func (iphone Iphone) call()  {

	fmt.Println("I am ipthon, I can call you ")

}



func  main()  {
	//切片

	var slice1 [] int =make([] int ,3)
	fmt.Println(slice1)

	slice2 :=  []int{1,2,4}
	fmt.Println(slice2)


	slice3 :=slice2[1:2]
	fmt.Println(slice3)

	fmt.Println(cap(slice1),len(slice1))


	var phone Phone

	phone = new(NokiaPthon)
	phone.call()


}



