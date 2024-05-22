package main

import (
	"fmt"
	"time"
)

func main() {
	go timer(12)
}

func timer(count int) {
	fmt.Println("Hey inside timer")
	for i := 0; i < count; i++ {
		fmt.Println(time.Now())
	}
	fmt.Println("Hey ending the timer")
}
