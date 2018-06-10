package main

import (
	"fmt"
	"os"
)

type LengthFunc func([]string) int
type LengthFuncFunc func(LengthFuncFunc) LengthFunc

func main() {
	y := func(makeLength func(LengthFunc) LengthFunc) LengthFunc {
		thing1 := func(recurser LengthFuncFunc) LengthFunc {
			return recurser(recurser)
		}
		thing2 := func(recurser LengthFuncFunc) LengthFunc {
			return makeLength(func(list []string) int {
				return recurser(recurser)(list)
			})
		}
		return thing1(thing2)
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
