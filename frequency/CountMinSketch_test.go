package frequency

import (
	"probabilistic-data-strutcures/skiplist/model"
	"testing"
)

func TestGetTheEstimateForASingleKey(t *testing.T) {
	countMinSketch := newCountMinSketch(10)

	key := model.NewSlice([]byte("Key"))
	countMinSketch.Increment(key)
	count := countMinSketch.Estimate(key)

	if count != 1 {
		t.Fatalf("Expected count to be 1, received %v", count)
	}
}

func TestGetTheEstimateForKeyWithMultipleOccurrences(t *testing.T) {
	countMinSketch := newCountMinSketch(10)

	key := model.NewSlice([]byte("Key"))
	countMinSketch.Increment(key)
	countMinSketch.Increment(key)
	countMinSketch.Increment(key)
	count := countMinSketch.Estimate(key)

	if count != 3 {
		t.Fatalf("Expected count to be 3, received %v", count)
	}
}

func TestGetTheEstimateForMultipleKeysWithMultipleOccurrences(t *testing.T) {
	countMinSketch := newCountMinSketch(10)

	key := model.NewSlice([]byte("Key"))
	otherKey := model.NewSlice([]byte("Other"))

	countMinSketch.Increment(key)
	countMinSketch.Increment(key)
	countMinSketch.Increment(key)

	countMinSketch.Increment(otherKey)

	count := countMinSketch.Estimate(key)
	if count != 3 {
		t.Fatalf("Expected count to be 3, received %v", count)
	}

	otherKeyCount := countMinSketch.Estimate(otherKey)
	if otherKeyCount != 1 {
		t.Fatalf("Expected otherKeyCount to be 1, received %v", otherKeyCount)
	}
}

func TestGetTheEstimateForKeysInStream(t *testing.T) {
	stream := []model.Slice{
		model.NewSlice([]byte("A")), model.NewSlice([]byte("B")),
		model.NewSlice([]byte("A")), model.NewSlice([]byte("C")),
		model.NewSlice([]byte("B")), model.NewSlice([]byte("A")),
		model.NewSlice([]byte("B")), model.NewSlice([]byte("C")),
	}
	expectedCounts := map[string]byte{
		"A": 3,
		"B": 3,
		"C": 2,
	}

	countMinSketch := newCountMinSketch(10)
	for _, key := range stream {
		countMinSketch.Increment(key)
	}

	for _, key := range stream {
		count := countMinSketch.Estimate(key)
		if count < expectedCounts[key.AsString()] {
			t.Fatalf("Expected atleast the count %v for key %v, received %v", expectedCounts[key.AsString()], key.AsString(), count)
		}
	}
}
