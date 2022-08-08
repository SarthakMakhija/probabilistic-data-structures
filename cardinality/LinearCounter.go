package cardinality

import (
	"github.com/spaolacci/murmur3"
	"math"
	"probabilistic-data-strutcures/skiplist/model"
	"unsafe"
)

type LinearCounter struct {
	byteVector    byteVector
	bitVectorSize int
}

func newLinearCounter(size int) *LinearCounter {
	byteVector := newByteVector(size)
	return &LinearCounter{
		byteVector:    byteVector,
		bitVectorSize: byteVector.bitSize(),
	}
}

func (linearCounter LinearCounter) Put(key model.Slice) {
	hash := murmur3.Sum32(key.GetRawContent())
	index := hash % uint32(linearCounter.bitVectorSize)
	bytePosition, mask := linearCounter.byteVector.bitPositionInByte(index)
	linearCounter.byteVector.setBitAt(bytePosition, mask)
}

func (linearCounter LinearCounter) Count() int {
	setBitCount := linearCounter.byteVector.countSetBits()
	if setBitCount == 0 || setBitCount == 1 || setBitCount == linearCounter.bitVectorSize {
		return setBitCount
	}
	estimation := math.Log(float64(linearCounter.bitVectorSize-setBitCount)) -
		math.Log(float64(linearCounter.bitVectorSize))

	return int(float64(linearCounter.bitVectorSize) * (-estimation))
}

var aByte byte

const byteSize = int(unsafe.Sizeof(&aByte))

type byteVector []byte

func newByteVector(size int) byteVector {
	return make(byteVector, size/byteSize+1)
}

func (bVector byteVector) bitSize() int {
	return len(bVector) * byteSize
}

func (bVector byteVector) setBitAt(position uint64, mask byte) {
	bVector[position] = bVector[position] | mask
}

func (bVector byteVector) countSetBits() int {
	countSetBits := func(n byte) int {
		count := 0
		for n != 0 {
			n = n & (n - 1)
			count = count + 1
		}
		return count
	}
	count := 0
	for index := 0; index < len(bVector); index++ {
		count = count + countSetBits(bVector[index])
	}
	return count
}

func (bVector byteVector) bitPositionInByte(keyIndex uint32) (uint64, byte) {
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
