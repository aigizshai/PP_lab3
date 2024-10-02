package main

import (
	ar "PP_LAB3/arrayutils"
	"PP_LAB3/mathutils"
	su "PP_LAB3/stringutils"
	"fmt"
)

func main() {
	fmt.Println("Введите число")
	var a int
	fmt.Scanln(&a)
	fmt.Println("Факториал числа ", mathutils.Factorial(a), "\n")

	fmt.Println("Введите строку")
	var str string
	fmt.Scanln(&str)
	fmt.Println("Перевернутая строка: ")
	res := su.Reverse(str)
	fmt.Println(res)

	var arr [5]int
	arr = ar.Create()
	ar.CreateSlice(arr[:])

	ar.CreateSliceStr()
}
