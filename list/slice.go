package list

import (
	"sync"
)

var _ List = (*Slice)(nil)

type Slice struct {
	items []*Item
}

func (s *Slice) Frequency(n uint) []Item {
	items := s.sortDesc(s.items)

	r := make([]Item, 0)
	for i, item := range items {
		if i >= int(n) {
			break
		}
		r = append(r, *item)
	}
	return r
}

func (s Slice) sortDesc(items []*Item) []*Item {
	l := len(items)

	if l == 0 {
		return items
	}
	pivot := items[0]

	left := make([]*Item, 0)
	right := make([]*Item, 0)
	for i := 1; i < l; i++ {
		if items[i].Count > pivot.Count {
			left = append(left, items[i])
			continue
		}
		right = append(right, items[i])
	}
	left = s.sortDesc(left)
	right = s.sortDesc(right)
	left = append(left, pivot)
	left = append(left, right...)

	return left
}

func (s Slice) sort(items []*Item) []*Item {
	l := len(items)

	if l == 0 {
		return items
	}
	pivot := items[0]

	left := make([]*Item, 0)
	right := make([]*Item, 0)
	for i := 1; i < l; i++ {
		if items[i].Count > pivot.Count {
			left = append(left, items[i])
			continue
		}
		right = append(right, items[i])
	}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		left = s.sort(left)
		wg.Done()
	}()
	go func() {
		right = s.sort(right)
		wg.Done()
	}()
	wg.Wait()

	left = append(left, pivot)
	left = append(left, right...)

	return left
}

func NewSlice() *Slice {
	return &Slice{}
}

func (s *Slice) Insert(str string) Item {
	for _, item := range s.items {
		if item.Name == str {
			item.Count++
			return *item
		}
	}
	item := &Item{Name: str, Count: 1}
	s.items = append(s.items, item)
	return *item
}
