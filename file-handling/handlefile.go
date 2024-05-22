package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//create file
	_, err := os.Create("timer.txt")
	if err != nil {
		return
	}
	wg.Add(1)
	fmt.Println("Created new file")
	fmt.Printf("this is time in string %v", time.Now().String())
	go timer(10)
	wg.Wait()

	msgChan := make(chan string)
	// channel impl
	go infiniteTimer(msgChan)

    //go putMethod()
    //go callMethod() 

	for {
		time, ok := <-msgChan
		if !ok {
			break
		}
		fmt.Printf("Receving data from channel!! time: %v\n", time)
	}
}

func timer(count int) {
	fmt.Printf("this is time in string %v", time.Now().String())
	var myTimeStr string
	for i := 0; i < count; i++ {
		myTimeStr = myTimeStr + "\n" + time.Now().String()
	}

	err := os.WriteFile("timer.txt", []byte(myTimeStr), 0666)
	if err != nil {
		return
	}
	fmt.Println(myTimeStr)
	wg.Done()
}

func infiniteTimer(message chan string) {
	for {
        message <- time.Now().String()
	}
}
