package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	(*b).color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PanitItBlack() {
	for i := range bl {
		(&bl[i]).SetColor(RED)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	bo1 := Box{
		1, 2, 3,
		RED,
	}
	bo2 := Box{
		1, 2, 3,
		RED,
	}
	bol := BoxList{
		bo1,
		bo2,
	}
	bol.PanitItBlack()
	fmt.Println(bol[1].color.String())

}
