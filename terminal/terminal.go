package terminal

import (
	"fmt"
)

type EsacpeSequence string

const (
	cursorHome   EsacpeSequence = "\033[H"
	eraseDisplay EsacpeSequence = "\033[2J"
	showCursor   EsacpeSequence = "\033[?25h"
	hideCursor   EsacpeSequence = "\033[?25l"
)

func CursorHome() {
	fmt.Printf("%v", cursorHome)
}

func EraseDisplay() {
	fmt.Printf("%v", eraseDisplay)
}

func ShowCursor() {
	fmt.Printf("%v", showCursor)
}

func HideCursor() {
	fmt.Printf("%v", hideCursor)
}
