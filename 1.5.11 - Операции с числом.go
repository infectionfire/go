/*Напишите программу, которая последовательно делает следующие операции с введённым числом:
Число умножается на 2;
Затем к числу прибавляется 100.*/

package main

import "fmt"

func main(){
	var a int
	var b int
	var c int
	fmt.Scan(&a)
	b = a * 2
	c = b + 100
	fmt.Println(c)

}