package main

import (
	"errors"
	"fmt"
	"time"
)

/*
	Facade — структурный паттерн проектирования, который предоставляет простой интерфейс к набору взаимосвязанных
	классов или объектов некоторой подсистемы, что облегчает ее использование.

	Use-cases:
	Когда нужно представить простой или урезанный интерфейс к сложной подсистеме
	Когда вы хотите разложить подсистему на отдельные слои

	Props:
	Изолирует клиентов от компонентов сложной подсистемы

	Cons:
	Фасад рискует стать божественным объектом, привязанным ко всем классам программы
*/

// Создадим все нужные структуры для работы системы
// Создадим банк
type Bank struct {
	Name  string
	Cards []Card
}

// Создадим карты для банка
type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

// СОздадим самого покупателя
type Buyer struct {
	Name string
	Card *Card
}

// создадим продукт
type Product struct {
	Name  string
	Price float64
}

// Создадим магазин
type Shop struct {
	Name    string
	Product []Product
}

// Создадим все функции работы самой бизнес логики
// Таймеры нужны для имитации работы и отсутствия мешанины в консоли
// Покупатель получает баланс своей карты
func (buyer Buyer) GetBalance() float64 {
	return buyer.Card.Balance
}

// Магазин продает товар
// Именно функция sell и является фасадом, так как в ней происходит основное взаимодействие со всеми остальными сервисами
func (shop Shop) Sell(buyer Buyer, product string) error {
	fmt.Println("[Магазин] Проверяем баланс покупателя")
	time.Sleep(time.Millisecond * 500)
	err := buyer.Card.CheckCardBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Баланс положительный. Проверяем возможность купить товар покупателю: %s\n", buyer.Name)
	time.Sleep(time.Millisecond * 500)
	for _, prod := range shop.Product {
		if prod.Name != product {
			continue
		}
		if prod.Price > buyer.GetBalance() {
			return errors.New("недостаточно средств для покупки")
		}
	}
	fmt.Printf("Товар: %s - успешно куплен", product)
	return nil
}

// Проверяем баланс
func (card Card) CheckCardBalance() error {
	fmt.Println("[Карта] Проверяем баланс в банке")
	time.Sleep(time.Millisecond * 500)
	return card.Bank.CheckBankBalance(card.Name)
}

// Проверяем баланс в банке
func (bank Bank) CheckBankBalance(cardNumber string) error {
	fmt.Printf("[Банк] Смотрим остаток на карте: %s\n", cardNumber)
	time.Sleep(time.Millisecond * 500)
	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("на карте недостаточно средств")
		}
	}
	fmt.Println("[Банк]На карте достаточно средств")
	return nil
}

/*
Итак в функциях выше была реализована многослойная логика когда пользователь хочет что-то купить,
Магазин запрашивает баланс у карты, карта смотрит у банка, все это дело возвращается, магазин проверяет наличие товара, проверяет может
ли покупатель его купить и пробает товар.
*/

// Инициализируем переменные для теста
var (
	bank = Bank{
		Name:  "bank",
		Cards: []Card{},
	}
	card1 = Card{
		Name:    "card-1",
		Balance: 100,
		Bank:    &bank,
	}
	card2 = Card{
		Name:    "card-2",
		Balance: 500,
		Bank:    &bank,
	}
	aboba = Buyer{
		Name: "Aboba",
		Card: &card1,
	}
	biba = Buyer{
		Name: "Biba",
		Card: &card2,
	}
	shop = Shop{
		Name: "Дукен",
		Product: []Product{
			prod,
		},
	}
	prod = Product{
		Name:  "Слива",
		Price: 250,
	}
)

func main() {
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("Банк выпустил карты для: %s, %s\n", aboba.Name, biba.Name)
	fmt.Printf("Покупатель - %s покупает товар: %s\n", aboba.Name, prod.Name)
	err := shop.Sell(aboba, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Покупатель - %s покупает товар: %s\n", biba.Name, prod.Name)
	err = shop.Sell(biba, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
