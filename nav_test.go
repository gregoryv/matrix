package position

import (
	"fmt"
)

func ExampleNavigator_Right() {
	nav := NewXYNavigator(0, 0, 0, 0, 2, 2)
	for x, y, more := nav.Right(); more; x, y, more = nav.Right() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
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
	nav := NewXYNavigator(1, 1, 0, 0, 1, 1)
	for x, y, more := nav.Up(); more; x, y, more = nav.Up() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 1,0
	// 0,1
	// 0,0
}

func ExampleNavigator_Left() {
	nav := NewXYNavigator(2, 2, 1, 1, 2, 2)
	for x, y, more := nav.Left(); more; x, y, more = nav.Left() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 1,2
	// 2,1
	// 1,1
}

func ExampleNavigator_Down() {
	nav := NewXYNavigator(1, 1, 1, 1, 2, 2)
	for x, y, more := nav.Down(); more; x, y, more = nav.Down() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 1,2
	// 2,1
	// 2,2
}
