package main

import "fmt"

func main() {
	fmt.Println(soma(112, 10) + divide(110, 5))
}

func soma(a int, b int) int {
	return a + b
}

func divide(a int, b int) int {
	return a / b
}
