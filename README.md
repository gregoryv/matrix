[![Build Status](https://travis-ci.org/gregoryv/matrix.svg?branch=master)](https://travis-ci.org/gregoryv/matrix)
[![codecov](https://codecov.io/gh/gregoryv/matrix/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/matrix)
[![Maintainability](https://api.codeclimate.com/v1/badges/6c81180a66486a8c24ab/maintainability)](https://codeclimate.com/github/gregoryv/matrix/maintainability)

[matrix](https://godoc.org/github.com/gregoryv/matrix) - Package implements iterator for matrix navigation


## Quick start

Use this package to minimize complexity when navigatin 2D matrix structures.
To loop over an matrix from top, left corner to bottom, right

	for nav, x, y, inside := NewNavigator(3, 3); inside; x, y, inside = nav.Right() {
		// do something
	}

Note that all navigation methods wrap at the edge of the given boundary.

	  0 1 2 3 4
	-+---------+
	0| | | | | |
	1| |-|-|-| |
	2| |-|>|x| |
	3| |x|x|x| |
	4| | | | | |
	-+---------+
