package bloom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitArray(t *testing.T) {

	for _, test := range []struct {
		//Bit to Set
		length    uint64
		bitsToSet []uint64
		negative  bool
	}{
		{length: 64, bitsToSet: []uint64{0, 10, 15, 16, 17, 20, 32, 43, 63}, negative: false},
		{length: 256, bitsToSet: []uint64{0, 10, 15, 16, 17, 20, 32, 43, 63, 64, 65, 72, 100, 102, 255}, negative: false},
		{length: 64 * 100000, bitsToSet: []uint64{0, 10, 15, 16, 17, 20, 32, 43, 63, 10000, 123466, 639999}, negative: false},
	} {
		ba := NewBitArray(test.length)
		m := make(map[uint64]bool)
		for _, i := range test.bitsToSet {
			ba.Set(i)
			m[i] = true
		}

		for i := uint64(0); i < test.length; i++ {
			if test.negative {
				assert.NotEqual(t, m[i], ba.IsSet(i))
			} else {
				assert.Equal(t, m[i], ba.IsSet(i))
			}
			if ba.IsSet(i) {
				ba.Clear(i)
				assert.NotEqual(t, m[i], ba.IsSet(i))
			}

		}
	}

}
