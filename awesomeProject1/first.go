package  main

import "fmt"

const MAX int = 3

type Book struct {
	author string
	subject string
	bookid int
}


func main ()  {
	var a  int =10
	fmt.Println(&a)
	//整形指针
	var ptr *int
	ptr =&a
	fmt.Println(*ptr)

	//结构体

	var book Book
	book.author="小明"
	fmt.Println(book)


	//结构体做为函数参数
	printBook(book)

	//结构体指针
	var struct_pointer *Book
	struct_pointer = &book
	fmt.Println(struct_pointer.author)

	var b []int


}

func printBook(book Book)  {
	book.subject="计算机"
	fmt.Println(book)
}
