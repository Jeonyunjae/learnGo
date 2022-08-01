package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}
	return total
}

func canIDrink(age int) bool {

	// 아래의 조건문과 같은 내용
	// koreanAge := age + 2
	// if koreanAge < 18

	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else {
		return true
	}
}

func canIDrinkSwitch(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	len, upper := lenAndUpper("nico")
	repeatMe("nico", "lynn", "jyj")
	fmt.Println(len, upper)

	fmt.Println(superAdd(1, 2, 3, 4, 5))
	fmt.Println(canIDrink(18))
}
