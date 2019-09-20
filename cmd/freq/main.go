package main

import (
	"bufio"
	"fmt"
	"freq"
	"github.com/pkg/errors"
	"log"
	"os"
	"time"
)

type iterator struct {
	file *os.File
	next chan []byte
	done chan bool
}

func newIterator(file *os.File) *iterator {

	it := &iterator{
		file: file,
		next: make(chan []byte),
		done: make(chan bool),
	}

	go func() {
		scanner := bufio.NewScanner(it.file)
		for scanner.Scan() {
			it.next <- scanner.Bytes()
		}
		it.done <- true
	}()

	return it
}

func (it *iterator) Next() (string, error) {
	select {
	case r := <-it.next:
		return string(r), nil
	case <-it.done:
		return "", errors.New("EOF")

	}
}

func main() {
	start := time.Now()
	if len(os.Args) == 1 {
		panic(errors.New("Filename was not received"))
	}
	name := os.Args[1]

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()
	defer func() { fmt.Printf("time execution:%v", time.Now().Sub(start)) }()

	it := newIterator(file)
	//counter := freq.NewSliceCounter()
	counter := freq.NewFileCounter()

	counter.ReadAll(it)

	for _, item := range counter.Frequency(20) {
		fmt.Printf("%d %s\n", item.Count, item.Name)
	}
}
