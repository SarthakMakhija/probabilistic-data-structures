package model

import (
	"probabilistic-data-strutcures/skiplist/comparator"
	"probabilistic-data-strutcures/skiplist/model"
	"probabilistic-data-strutcures/skiplist/utils"
	"testing"
)

func TestPutsAKeyValueAndGetByKeyInNode(t *testing.T) {
	const maxLevel = 8
	keyComparator := comparator.StringKeyComparator{}

	sentinelNode := NewNode(model.NilSlice(), model.NilSlice(), maxLevel)

	key := model.NewSlice([]byte("HDD"))
	value := model.NewSlice([]byte("Hard disk"))

	sentinelNode.Put(key, value, keyComparator, utils.NewLevelGenerator(maxLevel))

	getResult := sentinelNode.Get(key, keyComparator)
	if getResult.Value.AsString() != "Hard disk" {
		t.Fatalf("Expected %v, received %v", "Hard disk", getResult.Value.AsString())
	}
}

func TestPutAKeyValueAndAssertsItsExistenceInNode(t *testing.T) {
	const maxLevel = 8
	keyComparator := comparator.StringKeyComparator{}

	sentinelNode := NewNode(model.NilSlice(), model.NilSlice(), maxLevel)

	key := model.NewSlice([]byte("HDD"))
	value := model.NewSlice([]byte("Hard disk"))

	sentinelNode.Put(key, value, keyComparator, utils.NewLevelGenerator(maxLevel))

	getResult := sentinelNode.Get(key, keyComparator)
	if getResult.Exists != true {
		t.Fatalf("Expected key to exist, but it did not. Key was %v", "HDD")
	}
}

func TestPutAKeyValueAndAssertsItsNonExistenceInNode(t *testing.T) {
	const maxLevel = 8
	keyComparator := comparator.StringKeyComparator{}

	sentinelNode := NewNode(model.NilSlice(), model.NilSlice(), maxLevel)

	key := model.NewSlice([]byte("HDD"))
	value := model.NewSlice([]byte("Hard disk"))

	sentinelNode.Put(key, value, keyComparator, utils.NewLevelGenerator(maxLevel))

	getResult := sentinelNode.Get(model.NewSlice([]byte("NonExistent")), keyComparator)
	if getResult.Exists != false {
		t.Fatalf("Expected key to be missing, but it did exist. Key was %v", "NonExistent")
	}
}
