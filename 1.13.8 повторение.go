/*1.13.8
По данным числам, определите количество чисел, которые равны нулю.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество чисел, которые равны нулю.

Sample Input:

5
1
8
100
0
12
Sample Output:

1
 */


package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	slice := make([]int, n)
	var a int

	var j int = 0
	for fmt.Scan(&a); n >= 1 ; fmt.Scan(&a) {
		if j < n  {
			slice[j] = a
			j++
		}else{
			break
		}

	}
	num0 := 0
	for i := 0; i<n; i++{
		if slice[i] == 0{
			num0++
		}
	}
	fmt.Println(num0)
}