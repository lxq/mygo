// hello_2 project main.go
package main

import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	w, h, d float64
	color   Color
}

type BoxList []Box

func (box Box) Volume() float64 {
	return box.w * box.h * box.d
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestsColor() Color {
	k := Color(WHITE)
	v := 0.0
	for _, b := range bl {
		if b.Volume() > v {
			k = b.color
			v = b.Volume()
		}
	}
	return k
}

func (bl BoxList) PaintBlack() {
	for _, b := range bl {
		b.SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestsColor().String())
}
