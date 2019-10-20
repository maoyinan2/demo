package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, Go!")

	var mp = make(map[string]string)
	mp["a"] = "aaa"
	fmt.Println(mp)

}
