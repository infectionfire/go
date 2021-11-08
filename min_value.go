
package main

import "fmt"

func main() {
	var a, b, max int
	fmt.Scan(&a, &b)
	if a <= b {
		for i := a; i <= b; i++ {
			if i%7 == 0 {
				max = i
			}

		}
		if a == max || b == max {
			fmt.Println(max)
		} else if max == 0 {
			fmt.Println("NO")

		} else {
			fmt.Println(max)
		}

	}
}
