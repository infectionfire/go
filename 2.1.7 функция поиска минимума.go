/*2.1.7

Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:

4 5 6 7
Sample Output:

4
 */



package main
import "fmt"

func main()  {

	fmt.Println(minimumFromFour())


}
func minimumFromFour() int {

	var a int
	slice := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Scan(&a)
		slice[i] = a
	}
	y := slice[0]
	for i := 0; i < 4; i++ {
		if slice[i] < y {
			y = slice[i]
		}

	}
	return y
}