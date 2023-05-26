package main

import "fmt"

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}

/*Эта функция разделит значения int a и b.
Результат этого деления будет сохранен в int, на который указывает div.
Оставшаяся часть этого деления будет сохранена в int, указанном mod.*/
func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b

}
