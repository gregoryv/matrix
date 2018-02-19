[![Build Status](https://travis-ci.org/gregoryv/matrix.svg?branch=master)](https://travis-ci.org/gregoryv/matrix)
[![codecov](https://codecov.io/gh/gregoryv/matrix/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/matrix)


[matrix](https://godoc.org/github.com/gregoryv/position) - Package implements iterator for matrix navigation


## Quick start

Use this package to minimize complexity when navigatin 2D matrix structures.
To loop over an matrix from top, left corner to bottom, right

	x, y, inside := 1, 1, true
	nav := matrix.NewXYNavigator(x, y, 1, 1, 3, 3)
	for ; inside; x, y, inside = nav.Right() {
		// do something
	}


	  0 1 2 3 4
	-+---------+
	0| | | | | |
	1| |-|-|-| |
	2| |-|>|x| |
	3| |x|x|x| |
	4| | | | | |
	-+---------+
