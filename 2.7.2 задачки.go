/*2.7.2
На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

Sample Input:

6 8
Sample Output:

10
 */

package main

import "fmt"
import "math"

func main() {

	var a,b,c float64
	fmt.Scan(&a,&b)
	c = a*a+b*b
	c = math.Sqrt(c)
	fmt.Println(c)
}
