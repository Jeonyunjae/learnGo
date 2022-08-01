package main

import "fmt"

func main() {
	//슬라이스 갯수 제한이 필요없을때
	namesSlice := []string{"jyj", "kjh", "jjs", "jds"}
	namesSlice = append(namesSlice, "jdh")
	fmt.Println("Slice", namesSlice)

	//배열
	names := [5]string{"jyj", "kjh", "jjs", "jds"}
	fmt.Println("Array", names)

	//map
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println("Map", nico)

	for key, value := range nico {
		fmt.Println(key, value)
	}
}
