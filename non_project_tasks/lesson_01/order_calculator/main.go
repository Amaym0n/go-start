package main

// Урок 1. Задание 2: учебный калькулятор заказа.
//
// Это тренировочная CLI-программа, а не часть будущего production-сервиса.
// Она запрашивает:
//   - номинал карты;
//   - количество карт;
//   - промокод.
//
// Поддерживаемые номиналы и цены одной карты:
//
//	номинал 1000 -> цена 1090 рублей
//	номинал 2500 -> цена 2690 рублей
//	номинал 5000 -> цена 5290 рублей
//
// Правила:
//   - неподдерживаемый номинал нужно отклонить;
//   - количество должно быть от 1 до 10 включительно;
//   - промокод STARTGO дает скидку 100 рублей, если сумма до скидки
//     не меньше 5000 рублей;
//   - итоговая сумма не может быть отрицательной;
//   - ошибки ввода должны приводить к понятному сообщению;
//   - все денежные значения храни в целых рублях, без float32/float64.
//
// Обязательно выдели функции с именами:
//   - priceForNominal;
//   - calculateDiscount.
//
// Параметры и возвращаемые значения определи самостоятельно.
// priceForNominal должна позволять вызывающему коду понять,
// поддерживается ли введенный номинал.
//
// В конце выведи:
//   - номинал;
//   - количество;
//   - цену одной карты;
//   - сумму до скидки;
//   - размер скидки;
//   - итог к оплате.
//
// Сценарии для самостоятельной проверки:
//  1. Номинал 1000, количество 2, без промокода.
//  2. Номинал 2500, количество 2, промокод STARTGO.
//  3. Номинал 777, количество 1.
//  4. Поддерживаемый номинал, количество 0.
//  5. Вместо номинала введена строка.

import (
	"fmt"
)

type promocode struct {
	promo    string
	discount int
}

var startGo = promocode{"STARTGO", 100}

func priceForNominal(amount int) (int, bool) {
	switch amount {
	case 1000:
		return 1090, true
	case 2500:
		return 2690, true
	case 5000:
		return 5290, true
	default:
		return 0, false
	}
}

func calculateDiscount(amount int, promocode string) (int, int) {
	if promocode != startGo.promo {
		return 0, amount
	}
	if amount < 5000 {
		return 0, amount
	}
	newAmount := amount - startGo.discount
	return startGo.discount, newAmount
}

func main() {
	var amount, count int
	var promocode string
	// Номинал
	fmt.Println("Введите номинал 1000/2500/5000:")
	_, amountError := fmt.Scan(&amount)
	if amountError != nil {
		fmt.Printf("Некорректное значение, %s", amountError)
		return
	}
	price, priceBool := priceForNominal(amount)
	if !priceBool {
		fmt.Println("Недопустимый номинал")
		return
	}
	// Количество
	fmt.Println("Введите количество кодов от 1 до 10:")
	_, countError := fmt.Scan(&count)
	if countError != nil {
		fmt.Printf("Некорректное значение, %s", countError)
		return
	}
	if count < 1 || count > 10 {
		fmt.Println("Количество должно быть от 1 до 10")
		return
	}
	// Промокод
	fmt.Println("Введите промокод. Символ '-' будет являться пустой строкой:")
	_, promocodeError := fmt.Scan(&promocode)
	if promocodeError != nil {
		fmt.Printf("Некорректное значение, %s", promocodeError)
		return
	}
	totalAmount := price * count
	discount, newAmount := calculateDiscount(totalAmount, promocode)
	// Вывод
	fmt.Printf("Номинал: %d, количество: %d, цена одной карты: %d, сумма до скидки: %d, размер скидки: %d, итог к оплате: %d\n", amount, count, price, totalAmount, discount, newAmount)
}
