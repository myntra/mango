package bloom

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/willf/bloom"
	"math/rand"
	"testing"
)

func TestNewWithSize(t *testing.T) {
	bf := NewWithSize(61, 10)
	bf.Put("test1")
	bf.Put("test2")
	bf.Put("test3")
	assert.True(t, bf.MightContain("test1"))
	assert.True(t, bf.MightContain("test2"))
	assert.True(t, bf.MightContain("test3"))
	assert.False(t, bf.MightContain("test4"))
}

func TestBloomFilter(t *testing.T) {
	bf := New(1000, .03)

	bf.Put("test1")
	assert.True(t, bf.MightContain("test1"))

	assert.False(t, bf.MightContain("test2"))

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestBloomFilter_Put(t *testing.T) {
	count := 10000
	keys := make([]string, count)
	for i := 0; i < count; i++ {
		keys[i] = randStringRunes(50)
	}

	ukeys := make([]string, 2*count)
	for i := 0; i < 2*count; i++ {
		ukeys[i] = "A" + randStringRunes(50)
	}

	bf := New(uint64(count), .01)

	for i := 0; i < count; i++ {
		bf.Put(keys[i])
	}
	var fpc = 0
	for i := 0; i < count; i++ {
		assert.True(t, bf.MightContain(keys[i]))
	}
	for i := 0; i < 2*count; i++ {
		if bf.MightContain(ukeys[i]) {
			fpc++
		}
	}

	fmt.Printf("False Positive Count: %f", float32(fpc)/float32(count))

}

func BenchmarkBloomFilter_Put(b *testing.B) {
	count := b.N
	bf := New(uint64(count), .01)
	keys := make([]string, count)
	for i := 0; i < count; i++ {
		keys[i] = randStringRunes(50)
	}
	b.ResetTimer()
	for i := 0; i < count; i++ {
		bf.Put(keys[i])
	}
}

func BenchmarkBloomFilter_MightContain(b *testing.B) {
	count := b.N
	bf := New(uint64(count), .01)
	keys := make([]string, count)
	for i := 0; i < count; i++ {
		keys[i] = randStringRunes(50)
	}
	b.ResetTimer()
	for i := 0; i < count; i++ {
		bf.MightContain(keys[i])
	}
}

func BenchmarkBloomFilter_MightContain_Welly(b *testing.B) {
	count := b.N
	bf := bloom.NewWithEstimates(uint(count), float64(.01))
	keys := make([]string, count)
	for i := 0; i < count; i++ {
		keys[i] = randStringRunes(50)
	}
	b.ResetTimer()
	for i := 0; i < count; i++ {
		bf.TestString(keys[i])
	}
}

func BenchmarkBloomFilter_Add_Willy(b *testing.B) {
	count := b.N
	bf := bloom.NewWithEstimates(uint(count), float64(.01))
	keys := make([]string, count)
	for i := 0; i < count; i++ {
		keys[i] = randStringRunes(50)
	}
	b.ResetTimer()
	for i := 0; i < count; i++ {
		bf.AddString(keys[i])
	}
}
