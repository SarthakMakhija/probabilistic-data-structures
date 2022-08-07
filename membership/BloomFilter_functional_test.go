package membership

import (
	"probabilistic-data-strutcures/skiplist/model"
	"strconv"
	"testing"
)

func TestAdds500KeysAndChecksForTheirPositiveExistence(t *testing.T) {
	bloomFilter := newBloomFilter(5000, 0.1)

	keyUsing := func(count int) model.Slice {
		return model.NewSlice([]byte("Key-" + strconv.Itoa(count)))
	}
	for count := 1; count <= 500; count++ {
		_ = bloomFilter.Put(keyUsing(count))
	}

	for count := 1; count <= 500; count++ {
		contains := bloomFilter.Has(keyUsing(count))
		if contains == false {
			t.Fatalf("Expected key %v to be present but was not", keyUsing(count).AsString())
		}
	}
}

func TestAdds500KeysAndChecksForTheExistenceOfMissingKeys(t *testing.T) {
	bloomFilter := newBloomFilter(5000, 0.1)

	keyUsing := func(count int) model.Slice {
		return model.NewSlice([]byte("Key-" + strconv.Itoa(count)))
	}
	for count := 1; count <= 500; count++ {
		_ = bloomFilter.Put(keyUsing(count))
	}

	falsePositives := 0
	for count := 1; count <= 500; count++ {
		contains := bloomFilter.Has(keyUsing(count * 600))
		if contains == true {
			falsePositives = falsePositives + 1
		}
	}

	expectedFalsePositives := float64(bloomFilter.capacity) * bloomFilter.falsePositiveRate
	if float64(falsePositives) > expectedFalsePositives {
		t.Fatalf("Expected false positives %v, received %v", expectedFalsePositives, falsePositives)
	}
}
