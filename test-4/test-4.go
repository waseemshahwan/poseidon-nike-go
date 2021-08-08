package main

import (
	"bytes"
	"fmt"
)

func doSomethingAsync(data chan string) {}

func main() {

	fmt.Println([]byte{48,13,10,48})
	fmt.Println(bytes.Split([]byte{48,13,10,48}, []byte{13,10}))
}
