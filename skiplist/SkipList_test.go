package model

import (
	"probabilistic-data-strutcures/skiplist/comparator"
	"probabilistic-data-strutcures/skiplist/model"
	"testing"
)

func TestPutAKeyValueAndGetByKeyInSkiplist(t *testing.T) {
	memTable := NewSkipList(10, comparator.StringKeyComparator{})
	key := model.NewSlice([]byte("HDD"))
	value := model.NewSlice([]byte("Hard disk"))
	memTable.Put(key, value)

	getResult := memTable.Get(key)
	if getResult.Value.AsString() != "Hard disk" {
		t.Fatalf("Expected %v, received %v", "Hard disk", getResult.Value.AsString())
	}
}

func TestPutAKeyValueAndAssertsItsExistenceInSkiplist(t *testing.T) {
	memTable := NewSkipList(10, comparator.StringKeyComparator{})
	key := model.NewSlice([]byte("HDD"))
	value := model.NewSlice([]byte("Hard disk"))
	memTable.Put(key, value)

	getResult := memTable.Get(key)
	if getResult.Exists != true {
		t.Fatalf("Expected key to exist, but it did not. Key was %v", "HDD")
	}
}
