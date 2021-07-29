package main

import "fmt"

func main() {
	var a int
	a = 3
	var b *int
	fmt.Println("a=", a, "address", &a)
	fmt.Println("b=", b, "address", &b)
	b = &a
	fmt.Println("a=", a)
	fmt.Println("b=", *b)
	//обычная переменная
	//адресс(&)
	//значение

	//переменная ввиде ссылки(pointer)(*)
	//адресс
	//значение(*)
	//ссылку на другой элемент
}
