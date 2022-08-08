package cardinality

import (
	"probabilistic-data-strutcures/skiplist/model"
	"testing"
)

func TestCountWithDuplicates(t *testing.T) {
	linearCounter := newLinearCounter(10)
	linearCounter.Put(model.NewSlice([]byte("A")))
	linearCounter.Put(model.NewSlice([]byte("A")))
	linearCounter.Put(model.NewSlice([]byte("A")))
	linearCounter.Put(model.NewSlice([]byte("A")))
	linearCounter.Put(model.NewSlice([]byte("B")))
	linearCounter.Put(model.NewSlice([]byte("B")))
	linearCounter.Put(model.NewSlice([]byte("B")))
	linearCounter.Put(model.NewSlice([]byte("C")))
	linearCounter.Put(model.NewSlice([]byte("C")))
	linearCounter.Put(model.NewSlice([]byte("D")))

	count := linearCounter.Count()
	expected := 4

	if count != expected {
		t.Fatalf("Expected count to be %v, received %v", expected, count)
	}
}

func TestCountWithDistinct(t *testing.T) {
	linearCounter := newLinearCounter(10)
	linearCounter.Put(model.NewSlice([]byte("A")))
	linearCounter.Put(model.NewSlice([]byte("B")))
	linearCounter.Put(model.NewSlice([]byte("C")))
	linearCounter.Put(model.NewSlice([]byte("D")))

	count := linearCounter.Count()
	expected := 4

	if count != expected {
		t.Fatalf("Expected count to be %v, received %v", expected, count)
	}
}
