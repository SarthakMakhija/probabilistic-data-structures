package membership

import (
	"probabilistic-data-strutcures/skiplist/model"
	"testing"
)

func TestAddsAKeyWithBloomFilterAndChecksForItsPositiveExistence(t *testing.T) {
	bloomFilter := newBloomFilter(1, 0.001)

	key := model.NewSlice([]byte("Company"))
	_ = bloomFilter.Put(key)

	if bloomFilter.Has(key) == false {
		t.Fatalf("Expected %v key to be present but was not", key.AsString())
	}
}

func TestAddsAKeyWithBloomFilterAndChecksForTheExistenceOfANonExistingKey(t *testing.T) {
	bloomFilter := newBloomFilter(1, 0.001)

	key := model.NewSlice([]byte("Company"))
	_ = bloomFilter.Put(key)

	if bloomFilter.Has(model.NewSlice([]byte("Missing"))) == true {
		t.Fatalf("Expected %v key to be missing but was present", model.NewSlice([]byte("Missing")).AsString())
	}
}
