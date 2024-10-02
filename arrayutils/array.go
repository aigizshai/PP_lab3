package arrayutils

import (
	"fmt"
	"unicode/utf8"
)

func Create() [5]int {
	var a [5]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Println(a)
	return a
}

func CreateSlice(a []int) {
	var slice []int
	for i := 0; i < len(a); i++ {
		slice = append(slice, a[i])
	}

	fmt.Println("Исходный срез ", slice)

	slice = append(slice, 10, 20)
	fmt.Println("Срез после добавления ", slice)

	slice = append(slice[:2], slice[4:]...)
	fmt.Println("Срез после удаления ", slice)
}

func CreateSliceStr() string {
	s := make([]string, 0)
	fmt.Println("Введите 3 строки")
	for i := 0; i < 3; i++ {
		var temp string
		fmt.Scanln(&temp)
		s = append(s, temp)
	}
	fmt.Println("Введены ", s)
	len := 0
	var max string
	for i := 0; i < 3; i++ {
		if utf8.RuneCountInString(s[i]) > len {
			max = s[i]
			len = utf8.RuneCountInString(s[i])
		}
	}
	fmt.Println("Максимальная длина: ", max, " Длина: ", len)
	return max
}
