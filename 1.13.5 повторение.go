/*1.13.5
Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить, является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"

Sample Input:

6 8 10
Sample Output:

Прямоугольный
 */


package main

import "fmt"

func main() {

	array := [3]int{}
	var a int
	for i:=0; i < 3; i++{
		fmt.Scan(&a)
		array[i] = a
	}
	if (array[1]*array[1])+(array[0]*array[0])==array[2]*array[2]{
		fmt.Println("Прямоугольный")
	}else{
		fmt.Println("Непрямоугольный")
	}
}