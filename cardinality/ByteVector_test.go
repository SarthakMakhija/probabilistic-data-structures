package cardinality

import "testing"

func TestCountSetBitsInAByteVector(t *testing.T) {
	byteVector := newByteVector(10)
	byteVector.setBitAt(0, 0b00001000)
	byteVector.setBitAt(0, 0b00000001)
	byteVector.setBitAt(0, 0b00000010)

	countSetBits := byteVector.countSetBits()
	if countSetBits != 3 {
		t.Fatalf("Expected set bits to be 3, received %v", countSetBits)
	}
}
