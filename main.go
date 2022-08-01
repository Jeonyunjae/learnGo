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

func main() {
	len, upper := lenAndUpper("nico")
	repeatMe("nico", "lynn", "jyj")
	fmt.Println(len, upper)
}
