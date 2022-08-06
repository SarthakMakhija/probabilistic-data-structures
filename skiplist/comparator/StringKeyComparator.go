package comparator

import (
	"probabilistic-data-strutcures/skiplist/model"
	"strings"
)

type StringKeyComparator struct {
}

func (comparator StringKeyComparator) Compare(one model.Slice, other model.Slice) int {
	return strings.Compare(string(one.GetRawContent()), string(other.GetRawContent()))
}
