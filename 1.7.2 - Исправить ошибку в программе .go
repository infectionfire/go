/*Исправить ошибку в программе ниже:

Исходник:
package main

import "fmt"

func main(){
    var a int = 8
    const b int = 10
    a = a + b
    b = b + a
    fmt.Println(a)
}
 */

package main

import "fmt"

func main(){
	var a int = 8
	var b int = 10
	a = a + b
	b = b + a
	fmt.Println(a)
}