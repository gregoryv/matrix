// Package matrix implements 2D matrix navigation
package matrix

type XYNavigator struct {
	x, y int // current position
	edge Rect
}

type Rect struct {
	Left, Top     int // x,y of the top left corner
	Right, Bottom int // x,y of the bottom left corner
}

func NewNavigator(rows, cols int) (nav *XYNavigator, x, y int, inside bool) {
	x, y = 0, 0
	boundary := Rect{x, y, cols - 1, rows - 1}
	nav = NewXYNavigator(x, y, boundary)
	inside = true
	return
}

func NewXYNavigator(x, y int, boundary Rect) *XYNavigator {
	return &XYNavigator{x, y, boundary}
}

// Right returns next position right of the current and wraps lines until reaching xj, yj.
func (nav *XYNavigator) Right() (x, y int, inside bool) {
	nav.x++
	if nav.x > nav.edge.Right {
		// next row
		nav.x = nav.edge.Left
		nav.y++
	}
	return nav.x, nav.y, nav.y <= nav.edge.Bottom
}

// Left returns next position left of the current and wraps lines until reaching xi, yi.
func (nav *XYNavigator) Left() (x, y int, inside bool) {
	nav.x--
	if nav.x < nav.edge.Left {
		// previous row
		nav.x = nav.edge.Right
		nav.y--
	}
	return nav.x, nav.y, nav.y >= nav.edge.Top
}

// Up returns next position above the current and wraps columns until reaching xi, yi.
func (nav *XYNavigator) Up() (x, y int, inside bool) {
	nav.y--
	if nav.y < nav.edge.Top {
		// previous column
		nav.x--
		nav.y = nav.edge.Bottom
	}
	return nav.x, nav.y, nav.x >= nav.edge.Left
}

// Down returns next position below the current and wraps columns until reaching xj, yj.
func (nav *XYNavigator) Down() (x, y int, inside bool) {
	nav.y++
	if nav.y > nav.edge.Bottom {
		// next column
		nav.x++
		nav.y = nav.edge.Top
	}
	return nav.x, nav.y, nav.x <= nav.edge.Right
}
