//Заданы три числа - a, b, c ( a < b < c )
//a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить, является ли треугольник прямоугольным.
//Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"

package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	//Теорема пифагора
	//a*a+b*b=c*c
	s1 := a*a + b*b
	s2 := c * c

	if s1 == s2 {
		fmt.Print("Прямоугольный")
	} else {
		fmt.Print("Непрямоугольный")
	}
}
