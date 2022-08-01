package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"김치", "라면"}
	//nico := person{"nico", 18, favFood}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
}
