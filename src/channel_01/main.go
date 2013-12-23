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

func sum(a []int, c chan int) {
	t := 0
	for _, i := range a {
		t += i
	}
	c <- t //send value to channel
}

func main() {
	//go routine
	go sayHello("林秀全")
	sayHello("Hello, ")

	//channel 2013.12.18
	/*
		默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，
		这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。
		所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。
		其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。
		无缓冲channel是在多个goroutine之间同步很棒的工具。
	*/
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Printf("x:%d, y:%d, x+y:%d\n", x, y, x+y)

	//buffered channel 2013.12.19
	bc := make(chan int, 2) //修改2为1就报错，修改2为3可以正常运行
	bc <- 1
	bc <- 2
	fmt.Println(<-bc)
	fmt.Println(<-bc)

	//range

	//select

}
