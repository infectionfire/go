/*1.9.8 Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
Формат входных данных
На вход подается номер билета - одно шестизначное  число.
Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".
Sample Input:
613244
Sample Output:
YES
 */

package main

import "fmt"

func main() {
	var num int
	fmt.Print("")
	fmt.Scan(&num)
	var num1 int = num / 100000
	var num2 int = num / 10000 % 10
	var num3 int = num / 1000 %10
	var num4 int = num / 100 %10
	var num5 int = num / 10 %10
	var num6 int = num % 10
	var result1 = num1 + num2 + num3
	var result2 = num4 + num5 + num6
	if result1 == result2 {
		fmt.Println("YES")
	}else {
		fmt.Println("NO")		}
}
