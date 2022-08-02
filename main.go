package main

import (
	"fmt"
	"time"
)

go func main() {
	go sexyCount("yunjae")
	go sexyCount("jeonyou")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}

}
