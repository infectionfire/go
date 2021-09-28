package main

import "fmt"

func main() {
x := 10
p := &x

fmt.Println("значение х", x)
fmt.Println("адрес х",p)
fmt.Println("значение *p",*p)

*p = 15

fmt.Println("значение х после изменения *p", x)
fmt.Println("адрес х",p)
}