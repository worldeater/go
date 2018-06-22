package bitmap

import (
	"fmt"
	"math/rand"
)

const (
	bitsPerByte = 8
	allBitsSet  = 255
)

type Bitmap struct {
	data   []byte
	width  int
	height int
}

func New(width, height int) Bitmap {
	if (width < 0) || (height < 0) {
		panic(fmt.Sprintf("Bitmap size is negative: %v/%v", width, height))
	}

	// make our byte array big enough to hold all bits
	sizeInBytes := ((width * height) + bitsPerByte - 1) / bitsPerByte

	return Bitmap{
		data:   make([]byte, sizeInBytes),
		width:  width,
		height: height,
	}
}

func (bmap *Bitmap) getIndex(x, y int) int {
	return (x + bmap.width*y) / bitsPerByte
}

func (bmap *Bitmap) getOffset(x, y int) uint {
	return uint((x + bmap.width*y) % bitsPerByte)
}

func (bmap *Bitmap) Size() (width, height int) {
	return bmap.width, bmap.height
}

func (bmap *Bitmap) Get(x, y int) bool {
	index := bmap.getIndex(x, y)
	offset := bmap.getOffset(x, y)
	return ((bmap.data[index] >> offset) & 1) == 1
}

func (bmap *Bitmap) Set(x, y int) {
	index := bmap.getIndex(x, y)
	offset := bmap.getOffset(x, y)
	bmap.data[index] |= (1 << offset)
}

func (bmap *Bitmap) Clear(x, y int) {
	index := bmap.getIndex(x, y)
	offset := bmap.getOffset(x, y)
	bmap.data[index] &^= (1 << offset)
}

func (bmap *Bitmap) Toggle(x, y int) {
	index := bmap.getIndex(x, y)
	offset := bmap.getOffset(x, y)
	bmap.data[index] ^= (1 << offset)
}

func (bmap *Bitmap) ClearAll() {
	for i := range bmap.data {
		bmap.data[i] = 0
	}
}

func (bmap *Bitmap) SetAll() {
	for i := range bmap.data {
		bmap.data[i] = allBitsSet
	}
}

func (bmap *Bitmap) Randomize() {
	for i := range bmap.data {
		bmap.data[i] = byte(rand.Intn(255))
	}
}

func (bmap *Bitmap) RawData() *[]byte {
	return &bmap.data
}
