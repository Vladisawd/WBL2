package main

import "fmt"

/*
	Visitor - поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции,
	не изменяя классы объектов, над которыми эти операции могут выполняться.

	Use-cases:
	В проекте присутствуют объекты многих классов с различными интерфейсами
	 и нам необходимо выполнить над ними операции, которые зависят от конкретных классов;
	Необходимо выполнять не связанные между собой операции над объектами, которые
  	 входят в состав структуры и мы не хотим добавлять эти операции в классы;
	Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.

	Props:
	Упрощает добавление операций, работающих со сложными структурами объектов.
	Объединяет родственные операции в одном классе.
	Посетитель может накапливать состояние при обходе структуры элементов.

	Cons:
	Паттерн не оправдан, если иерархия элементов часто меняется (меняется логика визиторов)
	Может привести к нарушению инкапсуляции элементов.
*/

// Интерфейс элемента через который и будт новые операции
type Element interface {
	Accept(visitor Visitor)
}

// Сам паттерн по сути
type Visitor interface {
	VisitorElementA(ea *ElementA)
	VisitorElementB(eb *ElementB)
}

// Первый элемент
type ElementA struct{}

// Его метод
func (ea *ElementA) Accept(visitor Visitor) {
	visitor.VisitorElementA(ea)
}

// ВТорой
type ElementB struct{}

// Его
func (eb *ElementB) Accept(visitor Visitor) {
	visitor.VisitorElementB(eb)
}

type StrVisitor struct {
}

// Сами методы Элементов
func (sv *StrVisitor) VisitorElementA(elem *ElementA) {
	fmt.Println("Visiting Element A")
}

func (sv *StrVisitor) VisitorElementB(elem *ElementB) {
	fmt.Println("Visiting Element B")
}

func main() {
	visitor := &StrVisitor{}

	elementA := &ElementA{}
	elementA.Accept(visitor)

	elementB := &ElementB{}
	elementB.Accept(visitor)

}
