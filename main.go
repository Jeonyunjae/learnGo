package main

import "fmt"

func main() {
	//슬라이스 갯수 제한이 필요없을때
	namesSlice := []string{"jyj", "kjh", "jjs", "jds"}
	namesSlice = append(namesSlice, "jdh")
	fmt.Println(namesSlice)

	names := [5]string{"jyj", "kjh", "jjs", "jds"}
	fmt.Println(names)
}
