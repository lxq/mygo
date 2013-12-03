// channel_01 project main.go
package main

import (
	"fmt"
	"runtime"
)

func sayHello(str string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(str)
	}
}

func main() {
	fmt.Println("Go Routine!")
	go sayHello("林秀全")
	sayHello("Hello, ")
}
