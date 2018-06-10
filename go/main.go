package main

import "os"

type LengthFunc func([]string) int
type LengthFuncFunc func(LengthFuncFunc) LengthFunc

func main() {
	_ = func(makeLength func(LengthFunc) LengthFunc) LengthFunc {
		return func(recurser LengthFuncFunc) LengthFunc {
			return recurser(recurser)
		}(
			func(recurser LengthFuncFunc) LengthFunc {
				return makeLength(func(list []string) int {
					return recurser(recurser)(list)
				})
			},
		)
	}(
		func(length LengthFunc) LengthFunc {
			return func(list []string) int {
				if len(list) == 0 {
					return 0
				}
				return 1 + length(list[1:])
			}
		},
	)(os.Args[1:])
}
