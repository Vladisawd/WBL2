package main

import "fmt"

/*
	Chain of responsibility - поведенческий паттерн проектирования, который позволяет передавать запросы последовательно
	по цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

	Use-cases:
	Когда программа должна обрабатывать разнообразные запросы несколькими способами,
	 но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
	Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	Когда набор объектов, способных обработать запрос, должен задаваться динамически.

	Props:
	Уменьшает зависимость между клиентом и обработчиками.
	Реализует принцип единственной обязанности.
	Реализует принцип открытости/закрытости.
	Дополнительная гибкость при распределении обязанносте между объектами

	Cons:
	Запрос может остаться никем не обработанным.
*/

type Service interface {
	Exequte(*Data)
	SetNext(Service)
}

type Data struct {
	GetSource    bool //были би приняты данные
	UpdateSource bool // были ли обработаны ли данные
}

// Само устройство
type Device struct {
	Name string
	Next Service
}

// Устройство которые передает источник данных
// В этой функции мы проверяем не обработаны ли были уже ранее входные данные
func (device *Device) Exequte(data *Data) {
	if data.GetSource {
		fmt.Printf("Data: %s from device already get\n", device.Name)
		//если данные обработаны то сообщаем об этом и передаем следующему звену цепочки обработку
		device.Next.Exequte(data)
		return
	}
	//если нет, то ставим флаг что данные приняты и передаем следующему звену
	fmt.Printf("Get data: %s from device\n", device.Name)
	data.GetSource = true
	device.Next.Exequte(data)
}

// Это то самое следующее звено
func (device *Device) SetNext(srv Service) {
	device.Next = srv
}

// Поведение сервиса обработки данных
type DataUpdateService struct {
	Name string
	Next Service
}

// Сервис который обновляет полученные от устройства данные
func (update *DataUpdateService) Exequte(data *Data) {
	if data.GetSource {
		fmt.Printf("Data in service: %s is already update\n", update.Name)
		update.Next.Exequte(data)
		return
	}
	fmt.Printf("Ulready data: %s from service\n", update.Name)
	data.GetSource = true
	update.Next.Exequte(data)
}

func (update *DataUpdateService) SetNext(srv Service) {
	update.Next = srv
}

// Сервис который сохраняет полученные данные
type DataSaveService struct {
	Name string
	Next Service
}

// Сервис который обновляет полученные от устройства данные
func (save *DataSaveService) Exequte(data *Data) {
	if !data.GetSource {
		fmt.Printf("Data: %s not update\n", save.Name)
		return
	}
	fmt.Printf("Data: %s save\n", save.Name)
}

func (save *DataSaveService) SetNext(srv Service) {
	save.Next = srv
}

// Проверяем
func main() {
	//Создаем устройство
	device := &Device{
		Name: "device-1",
	}
	//Создаем проверку данных
	updateSrv := &DataUpdateService{
		Name: "updateSrv-1",
	}
	//Создаем сохранение данных
	saveData := &DataSaveService{
		Name: "saveData-1",
	}
	//Устройство передает данные на следующее звено - в проверку
	device.SetNext(updateSrv)
	//Проверка передает данные на следующее звено - в сохранение
	updateSrv.SetNext(saveData)

	//Сами данные
	data := &Data{}
	//запуск передачи
	device.Exequte(data)
}
