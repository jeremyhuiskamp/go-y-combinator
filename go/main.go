package main

type LengthFunc func([]string) int
type LengthFuncFunc func(LengthFuncFunc) LengthFunc

func main() {
	_ = func(makeLength func(LengthFunc) LengthFunc) LengthFunc {
		_ = func(recurser LengthFuncFunc) LengthFunc {
			return recurser(recurser)
		}
		_ = func(recurser LengthFuncFunc) LengthFunc {
			return makeLength(func(list []string) int {
				return recurser(recurser)(list)
			})
		}
		return nil
	}(
		func(length LengthFunc) LengthFunc {
			return func(list []string) int {
				if len(list) == 0 {
					return 0
				}
				return 1 + length(list[1:])
			}
		},
	)
}
