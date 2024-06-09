package main

import "fmt"

/*
	State - поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего
	состояния. Извне создаётся впечатление, что изменился класс объекта.
	Очень важным нюансом, отличающим этот паттерн от Стратегии, является то, что и контекст, и сами конкретные
	состояния могут знать друг о друге и инициировать переходы от одного состояния к другому.

	Use-cases:
	Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния,
	 причём типов состояний много, и их код часто меняется.
	Когда код класса содержит множество больших, похожих друг на друга, условных операторов, которые
	 выбирают поведения в зависимости от текущих значений полей класса.
	Когда вы сознательно используете табличную машину состояний, построенную на условных операторах,
	 но вынуждены мириться с дублированием кода для похожих состояний и переходов.

	Props:
	Избавляет от множества больших условных операторов машины состояний.
	Концентрирует в одном месте код, связанный с определённым состоянием.
	Упрощает код контекста.

	Cons:
	Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

// MobileAlertStater интерфейс для разных оповещений
type MobileAlertStater interface {
	Alert() string
}

// MobileAlert структура оповещений в зависимости от разных состояний
type MobileAlert struct {
	state MobileAlertStater
}

// функция возвращает само оповещение
func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// функция меняет состояние оповещения
func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// Конструктор для смены состояния
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertSong1{}}
}

// Одно из состояний
type MobileAlertSong1 struct {
}

// изменяем оповещение
func (a *MobileAlertSong1) Alert() string {
	return "Skibidi bibidi bi"
}

// Второе состояние
type MobileAlertSong struct {
}

// функция для второго состояния
func (a *MobileAlertSong) Alert() string {
	return "Never gonna give you up, Never gonna let you down"
}

func main() {
	//Инициализируем сами оповещения
	mobile := NewMobileAlert()

	//Вызываем их и выводим
	result := mobile.Alert()
	fmt.Println(result)

	//Пробуем вызвать несколько
	result += mobile.Alert()
	fmt.Println(result)

	//Добавляем еще одно
	mobile.SetState(&MobileAlertSong{})
	result = mobile.Alert()
	fmt.Println(result)
}
