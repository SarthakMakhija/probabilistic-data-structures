package comparator

import "probabilistic-data-strutcures/skiplist/model"

type KeyComparator interface {
	Compare(one model.Slice, other model.Slice) int
}
