package list

import (
	"fmt"
	"os"
	"path/filepath"
)

var _ List = (*File)(nil)

const cacheDir = "cache"

type File struct {
	dir string
}

func NewFile() *File {
	dir := filepath.Join(".", cacheDir)
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		fmt.Println(err, dir)
	}

	return &File{}
}

func (*File) Insert(s string) Item {
	//os.Open()
	item := &Item{s, 1}
	return *item
}

func (*File) Frequency(n uint) []Item {
	res := make([]Item, 0)
	return res
}
