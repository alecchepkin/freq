package freq

import (
	"freq/list"
	"log"
	"regexp"
	"strings"
)

type Iterator interface {
	Next() (string, error)
}

type Counter struct {
	list list.List
}

func NewSliceCounter() *Counter {
	c := &Counter{}
	c.list = list.NewSlice()
	return c
}

func NewFileCounter() *Counter {
	c := &Counter{}
	c.list = list.NewFile()
	return c
}

func (c *Counter) ReadAll(it Iterator) {
	for {
		row, err := it.Next()
		if err != nil {
			break
		}
		for _, s := range strings.Fields(row) {
			s = strings.ToLower(s)
			r, err := regexp.Compile("[^a-zA-Z]")
			if err != nil {
				log.Fatal(err)
			}
			s = r.ReplaceAllString(s, "")
			if len(s) == 0 {
				continue
			}
			c.list.Insert(s)
		}
	}
}

func (c *Counter) Frequency(n uint) []list.Item {
	return c.list.Frequency(n)
}
