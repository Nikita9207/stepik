//Цифровой корень
//Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр,
//на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации.
//Этот процесс повторяется до тех пор, пока не будет получена одна цифра.
//Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .
//По данному числу определите его цифровой корень.
//Входные данные
//Вводится одно натуральное число n, не превышающее 10^7
//Выходные данные
//Вывести цифровой корень числа n.

package main

import (
	"fmt"
)

func main() {
	var i, sum, num int

	fmt.Scan(&i)
	if i <= 0 || i > 100000000 {
		fmt.Print("Введите верное чилсло")
		return
	}

	sum = 0

	for i > 0 {
		sum += i % 10
		i /= 10
	}

	for sum > 0 {
		num += sum % 10
		sum /= 10
	}

	fmt.Print(num)
}