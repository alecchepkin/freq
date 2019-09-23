package list

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var _ List = (*File)(nil)

const cacheDir = "cache"
const resultFile = "result"

type File struct {
	dir string
}

func NewFile() *File {
	f := &File{}
	f.dir = filepath.Join(".", cacheDir)
	_ = os.RemoveAll(f.dir)
	_ = os.Mkdir(f.dir, os.ModePerm)
	fmt.Println("Reading...")
	return f

}

func (f *File) Insert(s string) Item {
	if len(s) == 0 {
		panic("string can not be empty")
	}
	dirs := strings.Join(strings.Split(s, ""), "/")
	path := filepath.Join(f.dir, dirs)
	err := os.MkdirAll(path, os.ModePerm)
	f.check(err)

	path = filepath.Join(path, "leaf")

	ff, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	f.check(err)
	defer ff.Close()

	buf := bytes.NewBuffer(nil)
	_, err = buf.ReadFrom(ff)
	f.check(err)

	item := &Item{Name: s}

	if buf.Len() > 0 {
		err = json.Unmarshal(buf.Bytes(), &item)
		f.check(err)
	}
	item.Count++
	res, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	_, err = ff.WriteAt(res, 0)
	f.check(err)

	return *item
}

func (File) check(err error) {
	if err != nil {
		panic(err)
	}
}

func (f *File) Frequency(n uint) []Item {
	fmt.Println("Finding frequency...")
	f.walkDirs()

	res := make([]Item, 0)

	for i := 0; i < int(n); i++ {
		max := f.nextMax(res)
		res = append(res, max)
	}
	return res
}

func (f *File) nextMax(res []Item) Item {
	var item Item
	max := Item{}
	ff, _ := os.Open(f.resultPath())
	defer ff.Close()
	scanner := bufio.NewScanner(ff)
	for scanner.Scan() {
		data := scanner.Bytes()
		err := json.Unmarshal(data, &item)
		f.check(err)
		if item.Count >= max.Count && !f.hasInRes(res, item.Name) {
			max = item
		}
	}
	return max
}

func (File) hasInRes(res []Item, name string) bool {
	for _, item := range res {
		if item.Name == name {
			return true
		}
	}
	return false
}

func (f *File) walkDirs() {
	ff, err := os.OpenFile(f.resultPath(), os.O_RDWR|os.O_CREATE, 0777)
	f.check(err)
	defer ff.Close()

	var wg sync.WaitGroup

	ch := make(chan []byte)

	go func() {
		for data := range ch {
			ff.Write(data)
			ff.WriteString("\n")
		}
	}()
	f.walkSubDir(f.dir, ch, &wg)
	wg.Wait()

}

func (f *File) walkSubDir(path string, ch chan<- []byte, wg *sync.WaitGroup) {

	wg.Add(1)
	go func(path string) {
		dirs, err := ioutil.ReadDir(path)
		f.check(err)
		for _, d := range dirs {
			subpath := filepath.Join(path, d.Name())
			if d.IsDir() {
				f.walkSubDir(subpath, ch, wg)
				continue
			}
			content, err := ioutil.ReadFile(subpath)
			f.check(err)
			ch <- content
		}
		wg.Done()
	}(path)

}

func (f *File) resultPath() string {
	return filepath.Join(f.dir, resultFile)
}
