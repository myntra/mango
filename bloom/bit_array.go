package bloom

import (
	"math"
)

type BitArray struct {
	length uint64
	array  []uint64
}

/**
Creates a New Bit Array of specified length
*/
func NewBitArray(length uint64) *BitArray {
	ba := &BitArray{length: length}
	ba.array = make([]uint64, int(math.Ceil(float64(length)/64)))
	return ba
}

/**
Sets bit in a particular index
*/
func (ba *BitArray) Set(index uint64) {
	ba.array[index/64] |= 1 << (index % 64)
}

/**
Clears bit in a particular index
*/
func (ba *BitArray) Clear(index uint64) {
	ba.array[index/64] &^= 1 << (index % 64)
}

/**
Returns true is bit is set in a particular index
*/
func (ba *BitArray) IsSet(index uint64) bool {
	return ((ba.array[index/64]) & (1 << (index % 64))) > 0
}
