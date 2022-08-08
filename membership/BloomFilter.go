package membership

import (
	"github.com/spaolacci/murmur3"
	"math"
	"probabilistic-data-strutcures/skiplist/model"
	"unsafe"
)

var aByte byte

const byteSize = int(unsafe.Sizeof(aByte))

type BloomFilter struct {
	capacity              int
	numberOfHashFunctions int
	falsePositiveRate     float64
	byteVector            []byte
}

func newBloomFilter(capacity int, falsePositiveRate float64) *BloomFilter {
	numberOfHashFunctions := numberOfHashFunctions(falsePositiveRate)
	bitVectorSize := bitVectorSize(capacity, falsePositiveRate)

	return &BloomFilter{
		capacity:              capacity,
		numberOfHashFunctions: numberOfHashFunctions,
		falsePositiveRate:     falsePositiveRate,
		byteVector:            make([]byte, bitVectorSize/byteSize+1),
	}
}

func (bloomFilter *BloomFilter) Put(key model.Slice) {
	indices := bloomFilter.keyIndices(key)
	for index := 0; index < len(indices); index++ {
		position, mask := bloomFilter.bitPositionInByte(indices[index])
		bloomFilter.byteVector[position] = bloomFilter.byteVector[position] | mask
	}
}

func (bloomFilter *BloomFilter) Has(key model.Slice) bool {
	indices := bloomFilter.keyIndices(key)
	for index := 0; index < len(indices); index++ {
		position, mask := bloomFilter.bitPositionInByte(indices[index])
		if bloomFilter.byteVector[position]&mask == 0 {
			return false
		}
	}
	return true
}

// Use the hash function to get all keyIndices of the given key
func (bloomFilter *BloomFilter) keyIndices(key model.Slice) []uint64 {
	indices := make([]uint64, 0, bloomFilter.numberOfHashFunctions)
	runHash := func(key []byte, seed uint32) uint64 {
		hash, _ := murmur3.Sum128WithSeed(key, seed)
		return hash
	}
	indexForHash := func(hash uint64) uint64 {
		return hash % uint64(bloomFilter.numberOfHashFunctions)
	}
	for index := 0; index < bloomFilter.numberOfHashFunctions; index++ {
		hash := runHash(key.GetRawContent(), uint32(index))
		indices = append(indices, indexForHash(hash))
	}
	return indices
}

func (bloomFilter *BloomFilter) bitPositionInByte(keyIndex uint64) (uint64, byte) {
	quotient, remainder := int64(keyIndex)/int64(byteSize), int64(keyIndex)%int64(byteSize)
	valueWithMostSignificantBit := int64(math.Pow(2, float64(byteSize)-1)) //128
	if remainder == 0 {
		return uint64(quotient), byte(valueWithMostSignificantBit)
	}
	return uint64(quotient), byte(0x0001 << (remainder - 1))
}
