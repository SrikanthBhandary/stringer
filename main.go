package main

import (
	"fmt"

	"github.com/srikanthbhandary/stringer"
	"github.com/srikanthbhandary/stringer/custom"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(stringer.MakeUpper("Test"))
	fmt.Println(stringer.Merger("Test1", "Test2"))
	fmt.Println(stringer.Merger("Test1", "Test2"))
	custom.Tests()
}
