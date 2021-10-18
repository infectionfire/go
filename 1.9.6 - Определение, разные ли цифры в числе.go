/*1.9.6 По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.
Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".
Sample Input 1:
237
Sample Output 1:
YES
Sample Input 2:
117
Sample Output 2:
NO
 */

package main

import "fmt"

func main() {
	var num int
	fmt.Print("")
	fmt.Scan(&num)
	var num1 int
	var num2 int
	var num3 int
	num1 = num / 100
	num2 = num / 10 %10
	num3 = num % 10

	if num1 != num2 && num2 != num3 && num1 != num3 {
				fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}