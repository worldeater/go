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
	width  uint
	height uint
}

func (b *Bitmap) getIndex(x, y uint) uint {
	return (x + b.width*y) / bitsPerByte
}

func (b *Bitmap) getOffset(x, y uint) uint {
	return (x + b.width*y) % bitsPerByte
}

func (b *Bitmap) checkBounds(x, y uint) {
	if (x > b.width-1) || (y > b.height-1) {
		panic(fmt.Sprintf("Bitmap index out of bounds: %v/%v", x, y))
	}
}

func (b *Bitmap) Size() (width, height uint) {
	return b.width, b.height
}

func (b *Bitmap) Init(width, height uint) {
	sizeInBytes := ((width * height) + bitsPerByte - 1) / bitsPerByte
	b.data = make([]byte, sizeInBytes)
	b.width = width
	b.height = height
}

func (b *Bitmap) Get(x, y uint) bool {
	b.checkBounds(x, y)
	index := b.getIndex(x, y)
	offset := b.getOffset(x, y)
	return ((b.data[index] >> offset) & 1) == 1
}

func (b *Bitmap) Set(x, y uint, bit bool) {
	b.checkBounds(x, y)
	index := b.getIndex(x, y)
	offset := b.getOffset(x, y)
	if bit {
		b.data[index] |= (1 << offset)
	} else {
		b.data[index] ^= (1 << offset)
	}
}

func (b *Bitmap) Flip(x, y uint) {
	b.checkBounds(x, y)
	index := b.getIndex(x, y)
	offset := b.getOffset(x, y)
	if ((b.data[index] >> offset) & 1) == 1 {
		b.data[index] ^= (1 << offset)
	} else {
		b.data[index] ^= (1 << offset)
	}
}

func (b *Bitmap) Clear() {
	for i := range b.data {
		b.data[i] = 0
	}
}

func (b *Bitmap) Fill() {
	for i := range b.data {
		b.data[i] = allBitsSet
	}
}

func (b *Bitmap) Randomize() {
	for i := range b.data {
		b.data[i] = byte(rand.Intn(255))
	}
}

func (b *Bitmap) Raw() *[]byte {
	return &b.data
}
