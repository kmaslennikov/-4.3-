package main

import "fmt"

func main() {
	nominalValue := 100
	maxCost := 100000
	var amount int

	fmt.Println("Банкомат.")
	fmt.Println("Введите сумму снятия со счёта:")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Операция не может быть выполнена.")
		fmt.Println("Сумма должна быть больше 0")
	} else if amount%nominalValue != 0 {
		fmt.Println("Операция не может быть выполнена.")
		fmt.Println("Банкомат выдаёт только купюры номиналом", nominalValue, "рублей")
	} else if amount > maxCost {
		fmt.Println("Операция не может быть выполнена.")
		fmt.Println("Сумма не может превышать", maxCost)
	} else {
		fmt.Println("Операция успешно выполнена.")
		fmt.Println("Вы сняли", amount, "рублей.")
	}
}
