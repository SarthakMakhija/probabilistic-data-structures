package membership

import (
	"github.com/bits-and-blooms/bitset"
	"github.com/spaolacci/murmur3"
	"probabilistic-data-strutcures/skiplist/model"
)

type BloomFilter struct {
	capacity              int
	numberOfHashFunctions int
	falsePositiveRate     float64
	bitVector             *bitset.BitSet
	bitVectorSize         uint
}

func newBloomFilter(capacity int, falsePositiveRate float64) *BloomFilter {
	numberOfHashFunctions := numberOfHashFunctions(falsePositiveRate)
	bitVectorSize := bitVectorSize(capacity, falsePositiveRate)

	return &BloomFilter{
		capacity:              capacity,
		numberOfHashFunctions: numberOfHashFunctions,
		falsePositiveRate:     falsePositiveRate,
		bitVector:             bitset.New(uint(bitVectorSize)),
		bitVectorSize:         uint(bitVectorSize),
	}
}

func (bloomFilter *BloomFilter) Put(key model.Slice) {
	indices := bloomFilter.keyIndices(key)
	for index := 0; index < len(indices); index++ {
		position := indices[index]
		bloomFilter.bitVector.Set(uint(position))
	}
}

func (bloomFilter *BloomFilter) Has(key model.Slice) bool {
	indices := bloomFilter.keyIndices(key)
	for index := 0; index < len(indices); index++ {
		position := indices[index]
		if !bloomFilter.bitVector.Test(uint(position)) {
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
		return hash % uint64(bloomFilter.bitVectorSize)
	}
	for index := 0; index < bloomFilter.numberOfHashFunctions; index++ {
		hash := runHash(key.GetRawContent(), uint32(index))
		indices = append(indices, indexForHash(hash))
	}
	return indices
}
