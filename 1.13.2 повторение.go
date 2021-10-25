/*1.13.2
Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:

745
Sample Output:

16
 */


package main


import "fmt"

func main() {

	var number int
	fmt.Scan(&number)
	var num1, num2, num3 int
	num1 = number / 100
	num2 = number / 10 % 10
	num3 = number %10
	fmt.Println(num1+num2+num3)

}