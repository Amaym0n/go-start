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
//
// TODO: реализуй решение внутри main и добавь необходимые функции.
import (
	"fmt"
)

type Promocode struct {
	promo    string
	discount int
}

var STARTGO = Promocode{"STARTGO", 100}

func priceForNominal(amount int) int {
	switch amount {
	case 1000:
		return 1090
	case 2500:
		return 2690
	case 5000:
		return 5290
	default:
		return 0
	}
}

func calculateDiscount(amount int, promocode string) (int, int, error) {
	if promocode != STARTGO.promo {
		return 0, amount, nil
	}
	if amount < 5000 {
		return 0, amount, nil
	}
	newAmount := amount - STARTGO.discount
	return STARTGO.discount, newAmount, nil
}

func main() {
	// TODO: твое решение.
	var amount, count int
	var promocode string
	// Номинал
	fmt.Println("Введите номинал 1000/2500/5000:")
	_, a_err := fmt.Scan(&amount)
	price := priceForNominal(amount)
	if price == 0 {
		fmt.Println("Недопустимый номинал")
		return
	}
	if a_err != nil {
		fmt.Printf("Некорректное значение, %s", a_err)
		return
	}
	// Количество
	fmt.Println("Введите количество кодов от 1 до 10:")
	_, c_err := fmt.Scan(&count)
	if c_err != nil {
		fmt.Printf("Некорректное значение, %s", c_err)
		return
	}
	if count < 1 || count > 10 {
		fmt.Println("Количество должно быть от 1 до 10")
		return
	}
	// Промокод
	fmt.Println("Введите промокод:")
	_, p_err := fmt.Scan(&promocode)
	if p_err != nil {
		fmt.Printf("Некорректное значение, %s", p_err)
		return
	}
	total_amount := price * count
	discount, newAmount, err := calculateDiscount(total_amount, promocode)
	if err != nil {
		fmt.Printf("Случилось непредвиденное исключение: %s", err)
		return
	}
	// Вывод
	fmt.Printf("Номинал: %d, количество: %d, цена одной карты: %d, сумма до скидки: %d, размер скидки: %d, итог к оплате: %d", amount, count, price, price*count, discount, newAmount)
}
