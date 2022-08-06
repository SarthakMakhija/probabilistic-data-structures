package model

import (
	"probabilistic-data-strutcures/skiplist/comparator"
	"probabilistic-data-strutcures/skiplist/model"
	"probabilistic-data-strutcures/skiplist/utils"
)

type SkipList struct {
	head           *Node
	keyComparator  comparator.KeyComparator
	levelGenerator utils.LevelGenerator
}

func NewSkipList(maxLevel int, keyComparator comparator.KeyComparator) *SkipList {
	return &SkipList{
		head:           NewNode(model.NilSlice(), model.NilSlice(), maxLevel),
		keyComparator:  keyComparator,
		levelGenerator: utils.NewLevelGenerator(maxLevel),
	}
}

func (skiplist *SkipList) Put(key, value model.Slice) bool {
	return skiplist.head.Put(key, value, skiplist.keyComparator, skiplist.levelGenerator)
}

func (skiplist *SkipList) Get(key model.Slice) model.GetResult {
	return skiplist.head.Get(key, skiplist.keyComparator)
}
