package terminal

import (
	"fmt"
)

func Home() {
	fmt.Printf("\033[H")
}

func Clear() {
	Home()
	fmt.Printf("\033[2J")
}

func HideCursor() {
	fmt.Printf("\033[?25l")
}

func ShowCursor() {
	fmt.Printf("\033[?25h")
}
