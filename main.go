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

func main() {
	len, upper := lenAndUpper("nico")
	repeatMe("nico", "lynn", "jyj")
	fmt.Println(len, upper)

	fmt.Println(superAdd(1, 2, 3, 4, 5))
}
