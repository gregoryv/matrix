// Package position implements iterator for matrix navigation
package position

type XYNavigator struct {
	x, y, xi, yi, xj, yj int
}

func NewXYNavigator(x, y, xi, yi, xj, yj int) *XYNavigator {
	return &XYNavigator{x, y, xi, yi, xj, yj}
}

// Right returns next position to the right, if the end of a row is reached
// following position starting at xi is returned.
func (nav *XYNavigator) Right() (x, y int, more bool) {
	nav.x++
	if nav.x > nav.xj {
		// next row
		nav.y++
		nav.x = nav.xi
	}
	if nav.y > nav.yj {
		return nav.x, nav.y, false
	}
	return nav.x, nav.y, true
}

func (nav *XYNavigator) Up() (x, y int, more bool) {
	nav.y--
	if nav.y < nav.yi {
		// previous column
		nav.x--
		nav.y = nav.yj
	}
	if nav.x < nav.xi {
		return nav.x, nav.y, false
	}
	return nav.x, nav.y, true
}

func (nav *XYNavigator) Down() (x, y int, more bool) {
	nav.y++
	if nav.y > nav.yj {
		// next column
		nav.x++
		nav.y = nav.yi
	}
	if nav.x > nav.xj {
		return nav.x, nav.y, false
	}
	return nav.x, nav.y, true
}

func (nav *XYNavigator) Left() (x, y int, more bool) {
	nav.x--
	if nav.x < nav.xi {
		// previous row
		nav.y--
		nav.x = nav.xj
	}
	if nav.y < nav.yi {
		return nav.x, nav.y, false
	}
	return nav.x, nav.y, true
}
