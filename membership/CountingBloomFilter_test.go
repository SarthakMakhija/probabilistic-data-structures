package membership

import (
	"probabilistic-data-strutcures/skiplist/model"
	"testing"
)

func TestAddsAKeyWithCountingBloomFilterAndChecksForItsPositiveExistence(t *testing.T) {
	bloomFilter := newCountingBloomFilter(1, 0.001)

	key := model.NewSlice([]byte("Company"))
	bloomFilter.Put(key)

	if bloomFilter.Has(key) == false {
		t.Fatalf("Expected %v key to be present but was not", key.AsString())
	}
}

func TestAddsAKeyWithCountingBloomFilterAndChecksForTheExistenceOfANonExistingKey(t *testing.T) {
	bloomFilter := newCountingBloomFilter(1, 0.001)

	key := model.NewSlice([]byte("Company"))
	bloomFilter.Put(key)

	if bloomFilter.Has(model.NewSlice([]byte("Missing"))) == true {
		t.Fatalf("Expected %v key to be missing but was present", model.NewSlice([]byte("Missing")).AsString())
	}
}

func TestAddsAndRemovesAKeyWithCountingBloomFilter(t *testing.T) {
	bloomFilter := newCountingBloomFilter(1, 0.001)

	key := model.NewSlice([]byte("Company"))
	bloomFilter.Put(key)
	bloomFilter.Remove(key)

	if bloomFilter.Has(key) == true {
		t.Fatalf("Expected %v key to be missing but was present", key.AsString())
	}
}

func TestAddsAndRemovesOneOfTheKeysWithCountingBloomFilter(t *testing.T) {
	bloomFilter := newCountingBloomFilter(20, 0.1)

	key1 := model.NewSlice([]byte("Cuckoo filter"))
	bloomFilter.Put(key1)

	key2 := model.NewSlice([]byte("X"))
	bloomFilter.Put(key2)

	bloomFilter.Remove(key1)

	if bloomFilter.Has(key2) == false {
		t.Fatalf("Expected %v key to be present but was missing", key2.AsString())
	}
	//False negative can happen in a Counting bloom filter
}
