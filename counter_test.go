package freq

import (
	"freq/list"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

type iterMock struct {
	slice []string
	cur   int
}

func newIterMock(slice []string) *iterMock {
	return &iterMock{slice: slice}
}

func (i *iterMock) Next() (string, error) {
	if i.cur < len(i.slice) {
		cur := i.cur
		i.cur++
		return i.slice[cur], nil
	}
	return "", errors.New("end")
}

func TestCounter_Frequency(t *testing.T) {

	tests := []struct {
		name string
		in   []string
		want []list.Item
	}{
		{"Simple", []string{"a"}, []list.Item{{"a", 1}}},
		{"should separate words", []string{"abc bcd"}, []list.Item{{"abc", 1}, {"bcd", 1}}},
		{"case insensitive", []string{"AbC"}, []list.Item{{"abc", 1}}},
		{"only letters", []string{"a!"}, []list.Item{{"a", 1}}},
	}
	for _, tt := range tests {
		c := NewSliceCounter()
		it := newIterMock(tt.in)
		t.Run(tt.name, func(t *testing.T) {
			c.ReadAll(it)
			if got := c.Frequency(20); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Counter.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
