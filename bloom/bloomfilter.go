package bloom

import (
	"github.com/spaolacci/murmur3"
	"math"
)

type BloomFilter struct {
	numOfHashFunctions int
	bitArray           *BitArray
}

/**
Create new BloomFilter with approximate unique keys insertion count and tolerated false positive probability
*/
func New(insertionCount uint64, falsePositiveProbability float64) *BloomFilter {
	length := optimalNumOfBits(insertionCount, falsePositiveProbability)
	hfs := optimalNumOfHashFunctions(insertionCount, length)

	bf := &BloomFilter{numOfHashFunctions: hfs, bitArray: NewBitArray(length)}
	return bf
}

/**
Create new BloomFilter with number of bits and hash functions to use
*/
func NewWithSize(numberOfBits uint64, hashFunctions int) *BloomFilter {
	return &BloomFilter{numOfHashFunctions: hashFunctions, bitArray: NewBitArray(numberOfBits)}
}

func optimalNumOfHashFunctions(insertions uint64, bitArrayLength uint64) int {
	return int(math.Max(float64(1), float64(bitArrayLength)/float64(insertions)*math.Log(float64(2))))
}

func optimalNumOfBits(insertions uint64, falsePositiveProbability float64) uint64 {
	return uint64(-float64(insertions) * math.Log(falsePositiveProbability) / (math.Log(float64(2)) * math.Log(float64(2))))
}

/**
Put the string key in bloom filter
*/
func (bf *BloomFilter) Put(key string) {
	bf.PutBytes([]byte(key))
}

/**
Put the bytes key in bloom filter
*/
func (bf *BloomFilter) PutBytes(key []byte) {
	bitLocations := bf.getBitLocationsToSet(key)
	for _, l := range bitLocations {
		bf.bitArray.Set(l)
	}
}

func (bf *BloomFilter) getBitLocationsToSet(key []byte) []uint64 {
	hasher := murmur3.New128()
	hasher.Write(key)
	h1, h2 := hasher.Sum128()
	bits := make([]uint64, bf.numOfHashFunctions)

	for i := 0; i < bf.numOfHashFunctions; i++ {
		combinedHash := h1 + uint64(i)*h2
		if combinedHash < 0 {
			combinedHash = ^combinedHash
		}

		bits[i] = combinedHash % bf.bitArray.length
	}
	return bits
}

/**
Check if bloom filter might contain string key
Can give false positives. But if returns false key is definitely not present
*/
func (bf *BloomFilter) MightContain(key string) bool {
	return bf.MightContainBytes([]byte(key))
}

/**
Check if bloom filter might contain bytes key
Can give false positives. But if returns false key is definitely not present
*/
func (bf *BloomFilter) MightContainBytes(key []byte) bool {
	bitLocations := bf.getBitLocationsToSet(key)
	for _, l := range bitLocations {
		if !bf.bitArray.IsSet(l) {
			return false
		}
	}
	return true
}
