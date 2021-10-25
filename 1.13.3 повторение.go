/*1.13.3
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.

Sample Input:

843
Sample Output:

348
 */


package main


import "fmt"

func main() {

	slice := make([]int, 3, 3)
	var n int
	fmt.Scan(&n)
	slice[0] = n%10
	slice[1] = n/10%10
	slice[2] = n/100
	fmt.Printf("%d%d%d",slice[0], slice[1],slice[2])

}