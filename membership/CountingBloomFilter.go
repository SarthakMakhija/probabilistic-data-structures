package membership

import (
	"github.com/spaolacci/murmur3"
	"math"
	"probabilistic-data-strutcures/skiplist/model"
	"unsafe"
)

var aByte byte

const byteSize = int(unsafe.Sizeof(&aByte))

type CountingBloomFilter struct {
	capacity              int
	numberOfHashFunctions int
	falsePositiveRate     float64
	byteVector            byteVector
	counterVector         counterVector
	bitVectorSize         uint
}

func newCountingBloomFilter(capacity int, falsePositiveRate float64) *CountingBloomFilter {
	numberOfHashFunctions := numberOfHashFunctions(falsePositiveRate)
	bitVectorSize := bitVectorSize(capacity, falsePositiveRate)
	byteVectorSize := bitVectorSize/byteSize + 1
	counterVectorSize := (byteVectorSize + (-byteVectorSize & (0x01))) / 2

	return &CountingBloomFilter{
		capacity:              capacity,
		numberOfHashFunctions: numberOfHashFunctions,
		falsePositiveRate:     falsePositiveRate,
		byteVector:            make(byteVector, byteVectorSize),
		counterVector:         make(counterVector, counterVectorSize),
		bitVectorSize:         uint(byteVectorSize * byteSize),
	}
}

func (bloomFilter *CountingBloomFilter) Put(key model.Slice) {
	indices := bloomFilter.keyIndices(key)
	for index := 0; index < len(indices); index++ {
		position, mask := bloomFilter.bitPositionInByte(indices[index])
		bloomFilter.counterVector.incrementBy1(position)
		bloomFilter.byteVector.set(position, mask)
	}
}

func (bloomFilter *CountingBloomFilter) Remove(key model.Slice) {
	exists, indices := bloomFilter.has(key)
	if !exists {
		return
	}
	for index := 0; index < len(indices); index++ {
		position, _ := bloomFilter.bitPositionInByte(indices[index])
		bloomFilter.counterVector.decrementBy1(position)
		if bloomFilter.counterVector.get(position) == 0 {
			bloomFilter.byteVector.clear(position)
		}
	}
}

func (bloomFilter *CountingBloomFilter) Has(key model.Slice) bool {
	exists, _ := bloomFilter.has(key)
	return exists
}

func (bloomFilter *CountingBloomFilter) has(key model.Slice) (bool, []uint64) {
	indices := bloomFilter.keyIndices(key)
	for index := 0; index < len(indices); index++ {
		position, mask := bloomFilter.bitPositionInByte(indices[index])
		if bloomFilter.byteVector.get(position)&mask == 0 {
			return false, indices
		}
	}
	return true, indices
}

// Use the hash function to get all keyIndices of the given key
func (bloomFilter *CountingBloomFilter) keyIndices(key model.Slice) []uint64 {
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

func (bloomFilter *CountingBloomFilter) bitPositionInByte(keyIndex uint64) (uint64, byte) {
	quotient, remainder := int64(keyIndex)/int64(byteSize), int64(keyIndex)%int64(byteSize)
	valueWithMostSignificantBit := int64(math.Pow(2, float64(byteSize)-1)) //128
	if remainder == 0 {
		if quotient == 0 {
			return uint64(quotient), byte(valueWithMostSignificantBit)
		}
		return uint64(quotient - 1), byte(valueWithMostSignificantBit)
	}
	return uint64(quotient), byte(0x01 << (remainder - 1))
}

type byteVector []byte
type counterVector []byte

func (b byteVector) clear(position uint64) {
	b[position] = 0
}

func (b byteVector) set(position uint64, mask byte) {
	b[position] = b[position] | mask
}

func (b byteVector) get(position uint64) byte {
	return b[position]
}

func (c counterVector) incrementBy1(position uint64) {
	index := position / 2
	shift := (position & 0x01) * 4
	isLessThan15 := (c[index]>>shift)&0x0f < 0x0f
	if isLessThan15 {
		c[index] = c[index] + (1 << shift)
	}
}

func (c counterVector) decrementBy1(position uint64) {
	index := position / 2
	if position&0x01 == 0x00 {
		//lower 4 bits: reduce by 1
		if (c[index] & 0x0f) == 0x00 {
			return
		}
		oneLess := (c[index] & 0x0f) - 1
		upperNibble := (c[index] >> 4) & 0x0f
		c[index] = (upperNibble << 4) | oneLess
	} else {
		//upper 4 bits: reduce by 1
		if (c[index] >> 4 & 0x0f) == 0x00 {
			return
		}
		lowerNibble := c[index] & 0x0f
		oneLess := ((c[index] >> 4) & 0x0f) - 1
		c[index] = (oneLess << 4) | lowerNibble
	}
}

func (c counterVector) get(position uint64) byte {
	index := position / 2
	shift := (position & 0x01) * 4
	return (c[index] >> shift) & 0x0f
}
