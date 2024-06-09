package main

import "fmt"

/*
	Command - поведенческий паттерн проектирования, который превращает запросы в объекты,
	позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их,
	а также поддерживать отмену операций.

	Use-cases:
	Когда вы хотите параметризовать объекты выполняемым действием.
	Когда вы хотите ставить операции в очередь, выполнять их по расписанию или передавать по сети.
	Когда вам нужна операция отмены

	Props:
	Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют
	Позволяет реализовать простую отмену и повтор операций
	Позволяет реализовать отложенный запуск операций
	Позволяет собирать сложные команды из простых
	Реализует принцип открытости/закрытости

	Cons:
	Усложняет код программы из-за введения множества дополнительных классов
*/

// Интерфейс команды
type Command interface {
	Execute()
}

// Структура с именем получателя
type Receiver struct {
	Name string
}

// Уведомление о том, что определенный получатель получил команду
func (r *Receiver) Action() {
	fmt.Println("Action has been taken by", r.Name)
}

// Сама команда
type ConcreteCommand struct {
	receiver *Receiver
}

// Встраиваем в структуру
func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Тот кто вызывает команду
type Invoker struct {
	command Command
}

// Вызов команды
func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

// Выход из нее
func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

// Все инициализируем и проверяем
func main() {
	receiver := &Receiver{Name: "Receiver1"}

	command := &ConcreteCommand{receiver: receiver}
	invoker := &Invoker{}

	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}
