package list

import (
	"reflect"
	"testing"
)

func TestFile_Insert(t *testing.T) {

	tests := []struct {
		name  string
		items []*Item
		in    string
		want  Item
	}{
		{"Insert first", []*Item{}, "a", Item{"a", 1}},
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
