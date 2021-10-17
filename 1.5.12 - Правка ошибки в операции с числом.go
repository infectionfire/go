/*Петя торопился в школу и неправильно написал программу,
которая сначала находит квадраты двух чисел, а затем их суммирует.
Исправьте его программу.

Исходный:

package main
import "fmt"
func main(){

  var a int
  fmt.Scan(&a) // считаем переменную 'a' с консоли
  fmt.Scan(&b) // считаем переменную 'b' с консоли

  a = a * a
  b = b * 2
  c = a + b
  fmt.Println(c)
}

*/

package main
import "fmt"
func main(){

	var a int
	var b int
	var c int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	a = a * a
	b = b * b
	c = a + b
	fmt.Println(c)
}