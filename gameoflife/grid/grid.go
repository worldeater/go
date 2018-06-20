package grid

import (
	"fmt"
	"github.com/worldeater/go/bitmap"
	"hash/crc32"
)

type Grid struct {
	bitmap.Bitmap
}

func (g *Grid) NextGen(nextGen *Grid) {
	nextGen.Clear()
	w, h := g.Size()
	var x, y uint
	for y = 0; y < h; y++ {
		for x = 0; x < w; x++ {
			n := g.AliveNeighbors(x, y)
			if n == 2 && g.GetBit(x, y) {
				nextGen.SetBit(x, y, true)
			}
			if n == 3 {
				nextGen.SetBit(x, y, true)
			}
		}
	}
}

func (g *Grid) ToString() string {
	const unicodeBrailleBase = 0x2800
	var (
		w, h   uint = g.Size()
		x, y   uint
		offset uint
		s      string
	)
	for y = 0; y < h; y += 4 {
		for x = 0; x < w; x += 2 {
			offset = 0
			if g.GetBit(x, y) {
				offset += 0x01
			}
			if g.GetBit(x, y+1) {
				offset += 0x02
			}
			if g.GetBit(x, y+2) {
				offset += 0x04
			}
			if g.GetBit(x+1, y) {
				offset += 0x08
			}
			if g.GetBit(x+1, y+1) {
				offset += 0x10
			}
			if g.GetBit(x+1, y+2) {
				offset += 0x20
			}
			if g.GetBit(x, y+3) {
				offset += 0x40
			}
			if g.GetBit(x+1, y+3) {
				offset += 0x80
			}
			s += fmt.Sprintf("%c", unicodeBrailleBase+offset)
		}
		s += fmt.Sprintln()
	}
	return s
}

func (g *Grid) AliveNeighbors(x2, y2 uint) int {
	n := 0
	x := int(x2)
	y := int(y2)
	// row above
	n += g.IsAlive(x-1, y-1)
	n += g.IsAlive(x-0, y-1)
	n += g.IsAlive(x+1, y-1)
	// left and right
	n += g.IsAlive(x-1, y-0)
	n += g.IsAlive(x+1, y-0)
	// row below
	n += g.IsAlive(x-1, y+1)
	n += g.IsAlive(x-0, y+1)
	n += g.IsAlive(x+1, y+1)
	return n
}

func (g *Grid) IsAlive(x2, y2 int) int {
	var (
		w, h uint = g.Size()
		x    uint = uint(x2)
		y    uint = uint(y2)
	)
	if x < 0 {
		x = w - 1
	}
	if y < 0 {
		y = h - 1
	}
	if g.GetBit(x%w, y%h) {
		return 1
	} else {
		return 0
	}
}

func (g *Grid) Checksum() uint32 {
	return crc32.ChecksumIEEE(*g.Raw())
}
