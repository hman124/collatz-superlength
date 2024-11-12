package main

import "fmt"

func main() {
	res := 2
	for i := 0; i < 3007; i++ {
		res = res * 2
		if res != 0 {
			fmt.Println(res)
		}
	}

	fmt.Println(res)
}
