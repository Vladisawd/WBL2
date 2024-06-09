package main

import (
	"fmt"
	"log"
)

/*
	Factory method - порождающий паттерн проектирования, который определяет общий интерфейс
	для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

	Use-cases:
	Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код
	Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки
	Когда вы хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых

	Props:
	Избавляет класс от привязки к конкретным классам продуктов
	Выделяет код производства продуктов в одно место, упрощая поддержку кода
	Упрощает добавление новых классов/структур в программу
	Реализует принцип открытости/закрытости

	Cons:
	Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя

*/

// Выведим в константы наши устройства
const (
	TypeComputer = "computer"
	TypePhone    = "phone"
	TypeTablet   = "tablet"
)

// Универсальный интерфейс для всех девайсов
type Device interface {
	GetTipe() string
	PrintDetails()
}

// Инициализация девайса
func NewDevice(typeName string) Device {
	switch typeName {
	case TypeComputer:
		return NewComputer()
	case TypePhone:
		return NewPhone()
	case TypeTablet:
		return NewTablet()
	default:
		log.Printf("Wrong type: %s\n", typeName)
		return nil
	}
}

// Сами девайсы и методы реализующие интерфейс для вывода типа и инфы
type Phone struct {
	Type           string
	ScreenDiagonal float32
	Memory         int
}

func NewPhone() Device {
	return Phone{
		Type:           TypePhone,
		ScreenDiagonal: 6.1,
		Memory:         256,
	}
}

func (p Phone) GetTipe() string {
	return p.Type
}

func (p Phone) PrintDetails() {
	fmt.Printf("Device information:\n Type: %s\n Screen diagonal: %.1f\n Memory: %d\n", p.Type, p.ScreenDiagonal, p.Memory)
}

type Computer struct {
	Type           string
	ScreenDiagonal float32
	Memory         int
}

func NewComputer() Device {
	return Phone{
		Type:           TypeComputer,
		ScreenDiagonal: 27.1,
		Memory:         2000,
	}
}

func (с Computer) GetTipe() string {
	return с.Type
}

func (с Computer) PrintDetails() {
	fmt.Printf("Device information:\n Type: %s\n Screen diagonal: %.1f\n Memory: %d\n", с.Type, с.ScreenDiagonal, с.Memory)
}

type Tablet struct {
	Type           string
	ScreenDiagonal float32
	Memory         int
}

func NewTablet() Device {
	return Phone{
		Type:           TypeComputer,
		ScreenDiagonal: 27.1,
		Memory:         2000,
	}
}

func (t Tablet) GetTipe() string {
	return t.Type
}

func (t Tablet) PrintDetails() {
	fmt.Printf("Device information:\n Type: %s\n Screen diagonal: %.1f\n Memory: %d\n", t.Type, t.ScreenDiagonal, t.Memory)
}

var types = []string{TypeComputer, TypePhone, TypeTablet, "Салфетка"}

// Проверка работы
func main() {
	for _, typeName := range types {
		device := NewDevice(typeName)
		if device == nil {
			continue
		}
		device.PrintDetails()
	}
}
