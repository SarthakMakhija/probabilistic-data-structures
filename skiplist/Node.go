package model

import (
	"probabilistic-data-strutcures/skiplist/comparator"
	"probabilistic-data-strutcures/skiplist/model"
	"probabilistic-data-strutcures/skiplist/utils"
)

type Node struct {
	key      model.Slice
	value    model.Slice
	forwards []*Node
}

func NewNode(key model.Slice, value model.Slice, level int) *Node {
	return &Node{
		key:      key,
		value:    value,
		forwards: make([]*Node, level),
	}
}

func (node *Node) Put(key model.Slice, value model.Slice, keyComparator comparator.KeyComparator, levelGenerator utils.LevelGenerator) bool {
	current := node
	positions := make([]*Node, len(node.forwards))

	for level := len(node.forwards) - 1; level >= 0; level-- {
		for current.forwards[level] != nil &&
			keyComparator.Compare(current.forwards[level].key, key) < 0 {
			current = current.forwards[level]
		}
		positions[level] = current
	}

	current = current.forwards[0]
	if current == nil || keyComparator.Compare(current.key, key) != 0 {
		newLevel := levelGenerator.Generate()
		newNode := NewNode(key, value, newLevel)
		for level := 0; level < newLevel; level++ {
			newNode.forwards[level] = positions[level].forwards[level]
			positions[level].forwards[level] = newNode
		}
		return true
	}
	return false
}

func (node *Node) Get(key model.Slice, keyComparator comparator.KeyComparator) model.GetResult {
	node, ok := node.nodeMatching(key, keyComparator)
	if ok {
		return model.GetResult{Key: key, Value: node.value, Exists: ok}
	}
	return model.GetResult{Key: key, Value: model.NilSlice(), Exists: false}
}

func (node *Node) nodeMatching(key model.Slice, keyComparator comparator.KeyComparator) (*Node, bool) {
	current := node
	for level := len(node.forwards) - 1; level >= 0; level-- {
		for current.forwards[level] != nil &&
			keyComparator.Compare(current.forwards[level].key, key) < 0 {
			current = current.forwards[level]
		}
	}
	current = current.forwards[0]
	if current != nil && keyComparator.Compare(current.key, key) == 0 {
		return current, true
	}
	return nil, false
}
