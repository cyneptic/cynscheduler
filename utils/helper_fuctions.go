package utils

import "fmt"

func Clear() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}
