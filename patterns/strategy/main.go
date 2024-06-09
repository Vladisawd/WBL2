package main

import "fmt"

/*
	Strategy - поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый
	 из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

	Use-cases:
	Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
	Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
	Когда вы не хотите обнажать детали реализации алгоритмов для других классов. (Стратегия позволяет изолировать код,
	 данные и зависимости алгоритмов от других объектов, скрыв эти детали внутри классов-стратегий)
	Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора.
	 Каждая ветка такого оператора представляет собой вариацию алгоритма.

	Props:
	Горячая замена алгоритмов на лету.
	Изолирует код и данные алгоритмов от остальных классов.
	Уход от наследования к делегированию.
	Реализует принцип открытости/закрытости.

	Cons:
	Усложняет программу за счёт дополнительных классов.
	Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

// Структура которая будет определять начальную и конечную точку пути
type Strategy interface {
	Route(startPoint int, endPoint int)
}

type Navigator struct {
	Strategy
}

func (nav *Navigator) SetStrategy(str Strategy) {
	nav.Strategy = str
}

// пример передвижения на машине
type RoadStrategy struct {
}

// ПОказываем всю инфу по перемещению на машине
func (r *RoadStrategy) Route(startPoint int, endPoint int) {
	speed := 40                          //скорость машины
	routeLength := endPoint - startPoint // длинна маршрута
	totalTime := routeLength * 40        //40 -среднее время
	fmt.Printf("Road A: %d, to B: %d\n speed: %d\n routeLength: %d\n totalTime: %d\n", startPoint, endPoint, speed, routeLength, totalTime)
}

// пример передвижения на автобусе
type BusStrategy struct {
}

// Показываем всю инфу по перемещению на автобусе
func (r *BusStrategy) Route(startPoint int, endPoint int) {
	speed := 30                          //скорость автобуса
	routeLength := endPoint - startPoint // длинна маршрута
	totalTime := routeLength * 50        //40 -среднее время
	fmt.Printf("Bus A: %d, to B: %d\n speed: %d\n routeLength: %d\n totalTime: %d\n", startPoint, endPoint, speed, routeLength, totalTime)
}

// пример передвижения пешком
type WalkStrategy struct {
}

// Показываем всю инфу по перемещению пешком
func (r *WalkStrategy) Route(startPoint int, endPoint int) {
	speed := 6                           //скорость пешая
	routeLength := endPoint - startPoint // длинна маршрута
	totalTime := routeLength * 60        //40 -среднее время
	fmt.Printf("Walk A: %d, to B: %d\n speed: %d\n routeLength: %d\n totalTime: %d\n", startPoint, endPoint, speed, routeLength, totalTime)
}

var (
	start      = 10
	end        = 40
	strategies = []Strategy{
		&RoadStrategy{},
		&BusStrategy{},
		&WalkStrategy{},
	}
)

// Видим рассчет одного и того же маршрута разными путями передвижения с разными данными, вот и паттерн стратегия!
func main() {
	nav := Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}
}
