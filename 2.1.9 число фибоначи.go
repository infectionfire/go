/*2.1.9

Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1. Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ... Напишите функцию, которая по указанному натуральному n возвращает φn.

Входные данные

Вводится одно число n.

Выходные данные

Необходимо вывести  значение φn.

Sample Input:

4
Sample Output:

3
 */

package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(fibonacci(number))
}



func fibonacci(n int) int {
	var table []int
	table = make([]int, n+1)
	table[0] = 0
	table[1] = 1


	for i := int(2); i <= n; i += 1 {

		table[i] = table[i-1] + table[i-2]


	}

	return table[n]
}