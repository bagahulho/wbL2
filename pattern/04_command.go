package main

import "fmt"

/*
"Команда" — поведенческий паттерн, который инкапсулирует запрос в виде объекта
Клиент через инвокер взаимодействует с получателем
Инвокер вызывает по нажатию кнопки определённую команду, которая знает своего получателя

Плюсы:
- Добавление новых команд без изменения кода
- Упрощает тестирование кода

Минусы:
- Усложнение кода за счёт увеличения количества объектов
- В некоторых случаях отмена команды может быть нерелевантной, что требует доп. обработки
*/

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Device interface {
	on()
	off()
}

type onCommand struct {
	device Device
}

func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device Device
}

func (c *offCommand) execute() {
	c.device.off()
}

type tv struct {
	isEnabled bool
}

func (t *tv) on() {
	t.isEnabled = true
	fmt.Println("tv is enabled")
}

func (t *tv) off() {
	t.isEnabled = false
	fmt.Println("tv is disabled")
}

type Command interface {
	execute()
}

func main() {
	tv := &tv{}
	on := &onCommand{
		device: tv,
	}
	off := &offCommand{
		device: tv,
	}

	onButton := &Button{
		command: on,
	}
	onButton.press()
	offButton := &Button{
		command: off,
	}
	offButton.press()
}
