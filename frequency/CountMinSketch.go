package frequency

import (
	"github.com/spaolacci/murmur3"
	"math/rand"
	"probabilistic-data-strutcures/skiplist/model"
	"time"
)

const depth = 4

type CountMinSketch struct {
	matrix [depth]row
	seeds  [depth]uint64
}

func newCountMinSketch(counters int) *CountMinSketch {
	next2Power := func(counters int64) int64 {
		counters--
		counters |= counters >> 1
		counters |= counters >> 2
		counters |= counters >> 4
		counters |= counters >> 8
		counters |= counters >> 16
		counters |= counters >> 32
		counters++
		return counters
	}

	source, updatedCounters := rand.New(rand.NewSource(time.Now().UnixNano())), next2Power(int64(counters))
	countMinSketch := &CountMinSketch{}

	for index := 0; index < depth; index++ {
		countMinSketch.seeds[index] = source.Uint64()
		countMinSketch.matrix[index] = make(row, updatedCounters/2)
	}
	return countMinSketch
}

func (countMinSketch *CountMinSketch) Increment(key model.Slice) {
	for index := 0; index < depth; index++ {
		hash := countMinSketch.runHash(key, uint32(countMinSketch.seeds[index]))
		currentRow := countMinSketch.matrix[index]
		currentRow.incrementAt(hash % depth)
	}
}

func (countMinSketch *CountMinSketch) Estimate(key model.Slice) byte {
	min := byte(255)
	for index := 0; index < depth; index++ {
		hash := countMinSketch.runHash(key, uint32(countMinSketch.seeds[index]))
		currentRow := countMinSketch.matrix[index]
		if valueAt := currentRow.getAt(hash % depth); valueAt < min {
			min = valueAt
		}
	}
	return min
}

func (countMinSketch *CountMinSketch) runHash(key model.Slice, seed uint32) uint64 {
	hash, _ := murmur3.Sum128WithSeed(key.GetRawContent(), seed)
	return hash
}

type row []byte

func (currentRow row) incrementAt(position uint64) {
	index := position / 2
	shift := (position & 0x01) * 4
	isLessThan15 := (currentRow[index]>>shift)&0x0f < 0x0f
	if isLessThan15 {
		currentRow[index] = currentRow[index] + (1 << shift)
	}
}

func (currentRow row) getAt(position uint64) byte {
	index := position / 2
	shift := (position & 0x01) * 4
	return (currentRow[index] >> shift) & 0x0f
}
