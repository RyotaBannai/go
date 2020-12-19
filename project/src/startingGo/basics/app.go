package main

import (
	"math"
	"fmt"
)

var (
	pkgVar = "Package Var"
)

func checkedOverflow(a, base uint32) bool {
	if (math.MaxUint32 - a) < base {
		print("The calculation causes " +
			"overflow and laps around the number.\n")
		return false
	} else {
		sum := base + a
		fmt.Printf("The result of calculation is %d\n", sum)
		return true
	}
}
