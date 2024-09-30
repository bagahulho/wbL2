package main

import "fmt"

/*
"Строитель" — порождающий паттерн проектирования.
Позволяет создавать сложные объекты, используя "шаги" —
— маленькие промежуточные объекты со своей простой логикой

Плюсы:
- Позволяет пошагово создать сложный продукт
- Позволяет использовать один и тот же код для создания различных объектов
- Изолирует сложную логику объекта

Минусы:
- Усложняет код из-за введения дополнительных объектов
- Привязка к конкретному объекту строителя
*/

type House struct {
	doorType   string
	windowType string
	floorCount int
}

type Builder interface {
	SetDoor()
	SetWindow()
	SetFloor()
}

type normalBuilder struct {
	doorType   string
	windowType string
	floorCount int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (nb *normalBuilder) SetDoor() {
	nb.doorType = "normalDoor"
}

func (nb *normalBuilder) SetWindow() {
	nb.windowType = "normalWindow"
}
func (nb *normalBuilder) SetFloor() {
	nb.floorCount = 2
}

func (nb *normalBuilder) getHouse() House {
	fmt.Println("BuildHouse")
	return House{
		doorType:   nb.doorType,
		windowType: nb.windowType,
		floorCount: nb.floorCount,
	}
}

func (nb *normalBuilder) BuildHouse() House {
	nb.SetDoor()
	nb.SetWindow()
	nb.SetFloor()
	return nb.getHouse()
}

func main() {
	builder := newNormalBuilder()
	newHouse := builder.BuildHouse()

	fmt.Printf("doorType in house: %s\n", newHouse.doorType)
	fmt.Printf("windowType in house: %s\n", newHouse.windowType)
	fmt.Printf("floorCount in house: %d\n", newHouse.floorCount)
}
