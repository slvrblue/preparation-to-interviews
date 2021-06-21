package main

import (
	"awesomeProject/compress"
	"fmt"
)

func main() {
	result, err := compress.Compress("AAAABBCDD")
	fmt.Println(result, err)

}
