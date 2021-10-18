/*1.10.4
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности,
во второй строке -- n чисел, входящих в данную последовательность.

Sample Input:

5
38 24 800 8 16
Sample Output:

40
 */

package main

import "fmt"

func main() {
	var sum, hoarder, n int
	fmt.Println("")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&hoarder)
		if hoarder%8 == 0 && hoarder < 100 && hoarder > 9 {
			sum = sum + hoarder
		}

	}
	fmt.Println(sum)
}