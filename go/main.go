package main

func main() {
	_ = func(list []string) int {
		if len(list) == 0 {
			return 0
		}
		return 1 + length(list[1:])
	}
}
