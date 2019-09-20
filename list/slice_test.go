package list

import (
	"reflect"
	"testing"
)

func TestSlice_Insert(t *testing.T) {

	tests := []struct {
		name  string
		items []*Item
		in    string
		want  Item
	}{
		{"Insert first", []*Item{}, "a", Item{"a", 1}},
		{"Insert second", []*Item{{"a", 1}}, "b", Item{"b", 1}},
		{"Insert duplicate", []*Item{{"a", 1}}, "a", Item{"a", 2}},
		{"case insensitive", []*Item{{"a", 1}}, "A", Item{"a", 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Slice{
				items: tt.items,
			}
			if got := s.Insert(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Frequency(t *testing.T) {

	tests := []struct {
		name  string
		items []*Item
		n     uint
		want  []Item
	}{
		{"one", []*Item{{"a", 1}}, 1, []Item{{"a", 1}}},
		{"desc", []*Item{{"b", 2}, {"a", 1}}, 2, []Item{{"b", 2}, {"a", 1}}},
		{"asc", []*Item{{"a", 1}, {"b", 2}}, 2, []Item{{"b", 2}, {"a", 1}}},
		{"asc", []*Item{{"a", 1}, {"b", 2}}, 2, []Item{{"b", 2}, {"a", 1}}},
		{"only n items", []*Item{{"a", 2}, {"c", 3}, {"b", 1}}, 2, []Item{{"c", 3}, {"a", 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Slice{
				items: tt.items,
			}
			if got := s.Frequency(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice.Frequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
