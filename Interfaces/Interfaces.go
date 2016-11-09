package main

import (
	"fmt"
	"strconv"
)

type myEnum int64

const (
	valueZero myEnum = iota
	valueOne
	valueTwo
)

func (m myEnum) String() string {
	return strconv.Itoa(int(m))
}

func main() {
	// x := myEnum(3)
	fmt.Printf("The value is : %s\n", valueOne)
}
