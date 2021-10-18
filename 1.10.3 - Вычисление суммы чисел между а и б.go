/*1.10.3
Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B
(каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.

Sample Input:

1 5
Sample Output:

15
 */

package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var sum int
	fmt.Println("")
	fmt.Scan(&num1)
	fmt.Println("")
	fmt.Scan(&num2)

	for ; num1 <= num2; num1++ {
		sum += num1
	}
	fmt.Println(sum)
}
