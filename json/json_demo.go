package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Person .
type Person struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"hobby,omitempty"`
}

// TagExample test.
type TagExample struct {
	Name         string `json:"name"`
	Ignored      int    `json:"-"`
	EmptyIgnored int    `json:"emptyIgnored,omitempty"`
	Field        int    `json:",omitempty"`
}

// Message .
type Message struct {
	Name, Text string
}

// Animal .
type Animal struct {
	Name  string `json:"name"`
	Order string `json:"order"`
}

func main() {
	// joe := Person{"joe", 20, []string{"Basketball", "Football", "Reading"}}
	// var b bytes.Buffer
	// n, _ := b.Write([]byte("Employee info: \n"))
	// fmt.Printf("Write %d bytes to buffer.\n", n)

	// jData, _ := json.Marshal(joe) // struct-》bytes.Buffer
	// fmt.Printf("%s\n", jData)

	// json.Compact(&b, jData) // 追加到bytes.Buffer
	// b.WriteTo(os.Stdout)

	var b bytes.Buffer
	myTeams := []Person{
		{"Joe", 18, []string{"<basketball>", "<football>", "<reading>"}},
		{"Jason", 20, []string{"pingpang&basketball", "Swim"}},
		{"Smith", 19, []string{"game&game"}},
	}
	_, err := b.Write([]byte(`<html><h1>Test<\h1><body>Run HTML Escape Example<script>`))
	jsonData, err := json.Marshal(myTeams)
	if err != nil {
		log.Fatal(err)
	}
	// <basketball>不想解释为tag就要转义，目前浏览器不帮你做
	json.HTMLEscape(&b, jsonData) // "hobby":["\u003cbasketball\u003e"
	_, err = b.Write([]byte(`</script></body></html>`))
	b.WriteTo(os.Stdout)

	var jsonBlob = []byte(`[
		{"name": "Platypus", "order": "Monotremata"},
		{"name": "Auoll", "order": "Dasyuromorphia"}
		]`)
	var animals = []Animal{}
	err = json.Unmarshal(jsonBlob, &animals) // bytes.Buffer-》struct
	fmt.Printf("\n%+v\n", animals)

}
