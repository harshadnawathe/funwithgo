package main

import (
	"fmt"

	"github.com/harshadnawathe/go-lang-fun/dsa/quickSort/partition"
)

func main() {
	a := []int{2, 1, 5, 4, 8, 6, 9, 3, 7}
	fmt.Println(a)
	p := partition.LomutoPartitionWithPivot(a, 0, 4)
	fmt.Println(a[0:p], a[p], a[p+1:])

}
