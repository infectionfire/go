/*1.9.7 Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:

1234
Sample Output:

1
 */

package main

import "fmt"

func main() {
	var num int
	fmt.Print("")
	fmt.Scan(&num)
	if num < 10 {
		fmt.Println(num)
	} else if num / 10 < 10 {
		fmt.Println(num / 10)
	} else if num / 100 < 10 {
		fmt.Println(num / 100)
	} else if num / 1000 < 10 {
		fmt.Println(num / 1000)
	} else if num / 10000 < 10 {
		fmt.Println(num / 10000)
	}else {
		fmt.Println("ERROR")		}
	}
