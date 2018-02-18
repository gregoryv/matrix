// Package position implements iterator for matrix navigation
package position

type XYNavigator struct {
	x, y   int // current position
	xi, yi int // top-left corner
	xj, yj int // bottom-right corner
}

func NewXYNavigator(x, y, xi, yi, xj, yj int) *XYNavigator {
	return &XYNavigator{x, y, xi, yi, xj, yj}
}

// Right returns next position right of the current and wraps lines until reaching xj, yj.
func (nav *XYNavigator) Right() (x, y int, inside bool) {
	nav.x++
	if nav.x > nav.xj {
		// next row
		nav.x = nav.xi
		nav.y++
	}
	return nav.x, nav.y, nav.y <= nav.yj
}

// Left returns next position left of the current and wraps lines until reaching xi, yi.
func (nav *XYNavigator) Left() (x, y int, inside bool) {
	nav.x--
	if nav.x < nav.xi {
		// previous row
		nav.x = nav.xj
		nav.y--
	}
	return nav.x, nav.y, nav.y >= nav.yi
}

// Up returns next position above the current and wraps columns until reaching xi, yi.
func (nav *XYNavigator) Up() (x, y int, inside bool) {
	nav.y--
	if nav.y < nav.yi {
		// previous column
		nav.x--
		nav.y = nav.yj
	}
	return nav.x, nav.y, nav.x >= nav.xi
}

// Down returns next position below the current and wraps columns until reaching xj, yj.
func (nav *XYNavigator) Down() (x, y int, inside bool) {
	nav.y++
	if nav.y > nav.yj {
		// next column
		nav.x++
		nav.y = nav.yi
	}
	return nav.x, nav.y, nav.x <= nav.xj
}
