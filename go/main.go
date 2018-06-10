package main

import (
	"fmt"
	"os"
)

type LengthFunc func([]string) int
type LengthFuncFunc func(LengthFuncFunc) LengthFunc

func main() {
	applyToSelf := func(recurser LengthFuncFunc) LengthFunc {
		return recurser(recurser)
	}
	y := func(makeLength func(LengthFunc) LengthFunc) LengthFunc {
		thing2 := func(recurser LengthFuncFunc) LengthFunc {
			return makeLength(func(list []string) int {
				return recurser(recurser)(list)
			})
		}
		return applyToSelf(thing2)
	}
	incompleteLength := func(length LengthFunc) LengthFunc {
		return func(list []string) int {
			if len(list) == 0 {
				return 0
			}
			return 1 + length(list[1:])
		}
	}
	length := y(incompleteLength)

	fmt.Println(length(os.Args[1:]))
}
