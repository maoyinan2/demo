package main

import (
	"fmt"
	"sync"
)

type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func main() {
	s := New()
	s["this"] = "that"
	fmt.Println(s)
	s2 := New()
	fmt.Println(s2)

}

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}
