package main

import (
	"fmt"
	"github.com/kumakichi/wol"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s mac1 [mac2] ...\n", os.Args[0])
		return
	}

	for i := 1; i < len(os.Args); i++ {
		err := wol.Wake(os.Args[i])
		if err != nil {
			fmt.Printf("error Wake, err: %v\n", err)
			return
		}
	}
}
