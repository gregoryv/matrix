package matrix

import (
	"fmt"
)

func ExampleNavigator_Right() {
	x, y, inside := 0, 0, true
	boundary := Rect{x, y, 2, 2}
	nav := NewXYNavigator(x, y, boundary)
	for ; inside; x, y, inside = nav.Right() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 0,0
	// 1,0
	// 2,0
	// 0,1
	// 1,1
	// 2,1
	// 0,2
	// 1,2
	// 2,2
}

func ExampleNavigator_Right_full() {
	x, y, inside := 0, 0, true
	boundary := Rect{x, y, 2, 2}
	nav := NewXYNavigator(x, y, boundary)
	for ; inside; x, y, inside = nav.Right() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 0,0
	// 1,0
	// 2,0
	// 0,1
	// 1,1
	// 2,1
	// 0,2
	// 1,2
	// 2,2
}

func ExampleNavigator_Up() {
	x, y, inside := 1, 1, true
	boundary := Rect{0, 0, x, y}
	nav := NewXYNavigator(x, y, boundary)
	for ; inside; x, y, inside = nav.Up() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 1,1
	// 1,0
	// 0,1
	// 0,0
}

func ExampleNavigator_Left() {
	x, y := 2, 2
	boundary := Rect{1, 1, x, y}
	nav := NewXYNavigator(x, y, boundary)
	for x, y, more := nav.Left(); more; x, y, more = nav.Left() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 1,2
	// 2,1
	// 1,1
}

func ExampleNavigator_Down() {
	x, y, inside := 1, 1, true
	boundary := Rect{x, y, 2, 2}
	nav := NewXYNavigator(x, y, boundary)
	for ; inside; x, y, inside = nav.Down() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 1,1
	// 1,2
	// 2,1
	// 2,2
}
