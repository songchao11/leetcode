package main

import (
	"fmt"
)

//3211
//i:1    1   2    3
//x:321  32  3    0
//r:1    11  112  1123
func main() {
	x := 1210
	if x < 0 {
	}
	k := x
	res := 0
	for k != 0 {
		i := k % 10
		k = k / 10
		res = res*10 + i
	}
	r := x == res
	fmt.Println(r)
}
