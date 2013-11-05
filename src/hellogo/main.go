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

func main() {
	fmt.Println("Hello World! 你好，世界！")
	fmt.Printf("2的开根号：%v \n", mymath.Sqrt(2))
	
	r1 := Rect{123, 456}
	fmt.Printf("面积是：%d\n", r1.Area())
}
