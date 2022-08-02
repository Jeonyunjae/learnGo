package main

import (
	"fmt"
	"name/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	err := dictionary.Update(baseWord, "Sec")
	if err != nil {
		fmt.Println(err)
	}
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
