// hellogo project main.go
package main

import (
	"fmt"
	"mymath"
)

type Rect struct {
	w, h int
}

func (rc Rect) Area() int {
	return rc.w * rc.h
}

func sum(a []int, c chan int) {
	t := 0
	for _, v := range a {
		t += v
	}
	c <- t //send total to c
}

func main() {
	fmt.Println("Hello World! 你好，世界！")
	fmt.Printf("2的开根号：%v \n", mymath.Sqrt(2))
	
	r1 := Rect{123, 456}
	fmt.Printf("面积是：%d\n", r1.Area())
	
	a := []int{7, 2, 8, -9, 4, 0}
	c := make (chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c //receive from c
	fmt.Printf("total = %d\n", x + y)
}
