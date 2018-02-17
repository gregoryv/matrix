[![Build Status](https://travis-ci.org/gregoryv/position.svg?branch=master)](https://travis-ci.org/gregoryv/position)
[![codecov](https://codecov.io/gh/gregoryv/position/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/position)


[position](https://godoc.org/github.com/gregoryv/position) - Package implements iterator for matrix navigation


## Quick start

Use this package to minimize complexity when navigatin 2D matrix structures.
To loop over an matrix from top, left corner to bottom, right

	// Start outside the matrix
	x, y := -1, 0
	nav := NewXYNavigator(x, y, 0, 0, 2, 2)
	for x, y, more := nav.Right(); more; x, y, more = nav.Right() {
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
