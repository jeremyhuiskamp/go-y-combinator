package main

import (
	"fmt"
	"os"
)

type LengthFunc func([]string) int
type LengthFuncFunc func(LengthFuncFunc) LengthFunc

func main() {
	y := func(makeLength func(LengthFunc) LengthFunc) LengthFunc {
		recurser := func(recurser LengthFuncFunc) LengthFunc {
			return makeLength(func(list []string) int {
				return recurser(recurser)(list)
			})
		}
		return recurser(recurser)
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
