package main

import (
	"fmt"
	"os"
)

func main() {
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Println(os.Args[i])
	// }
	// for i, v := range os.Args {
	// 	if i != 0 {
	// 		fmt.Println(v)
	// 	}
	// }
	for _, v := range os.Args[1:] {
		fmt.Println(v)
	}
}
