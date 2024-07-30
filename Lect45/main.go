// Modules

package main

import (
	"fmt"

	"github.com/dogeyboy1932/goPuppy"
)

func main() {
	s1 := goPuppy.Bark()
	s2 := goPuppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := goPuppy.BigBark()
	s4 := goPuppy.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)
}
