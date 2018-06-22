package grid

import (
	"fmt"
	"github.com/worldeater/go/bitmap"
)

type Grid struct {
	bitmap.Bitmap
}

func New(width, height int) Grid {
	return Grid{bitmap.New(width, height)}
}

func (grid *Grid) NextGen(nextGen *Grid) {
	nextGen.ClearAll()
	w, h := grid.Size()
	var x, y int
	for y = 0; y < h; y++ {
		for x = 0; x < w; x++ {
			n := grid.CountNeighbors(x, y)
			if n == 2 && grid.Get(x, y) {
				nextGen.Set(x, y)
			}
			if n == 3 {
				nextGen.Set(x, y)
			}
		}
	}
}

func (grid *Grid) ToString() string {
	// see https://en.wikipedia.org/wiki/Braille_Patterns#Identifying,_naming_and_ordering
	const (
		unicodeBrailleBase = 0x2800
		brailleWidth       = 2
		brailleHeight      = 4
	)
	var (
		w, h int = grid.Size()
		s    string
	)
	for y := 0; y < h; y += brailleHeight {
		for x := 0; x < w; x += brailleWidth {
			offset := 0
			if grid.Get(x, y) {
				offset += 0x01
			}
			if grid.Get(x, y+1) {
				offset += 0x02
			}
			if grid.Get(x, y+2) {
				offset += 0x04
			}
			if grid.Get(x+1, y) {
				offset += 0x08
			}
			if grid.Get(x+1, y+1) {
				offset += 0x10
			}
			if grid.Get(x+1, y+2) {
				offset += 0x20
			}
			if grid.Get(x, y+3) {
				offset += 0x40
			}
			if grid.Get(x+1, y+3) {
				offset += 0x80
			}
			s += fmt.Sprintf("%c", unicodeBrailleBase+offset)
		}
		s += fmt.Sprintln()
	}
	return s
}

func (grid *Grid) Alive(x, y int) int {
	var w, h int = grid.Size()
	if x < 0 {
		x = w - 1
	} else if x >= w {
		x -= w
	}
	if y < 0 {
		y = h - 1
	} else if y >= h {
		y -= h
	}
	if grid.Get(x, y) {
		return 1
	} else {
		return 0
	}
}

func (grid *Grid) CountNeighbors(x, y int) int {
	n := 0
	// row above
	n += grid.Alive(x-1, y-1)
	n += grid.Alive(x-0, y-1)
	n += grid.Alive(x+1, y-1)
	// left and right
	n += grid.Alive(x-1, y-0)
	n += grid.Alive(x+1, y-0)
	// row below
	n += grid.Alive(x-1, y+1)
	n += grid.Alive(x-0, y+1)
	n += grid.Alive(x+1, y+1)
	return n
}
