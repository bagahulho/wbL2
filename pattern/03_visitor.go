package main

import (
	"fmt"
	"math"
)

/*
"Посетитель" добавляет новый функционал объекту без его изменения.

Плюсы:
- не вносим в структуру объекта множество не свяазанных между собой операции
- не изменяем структуру объекта
- спокойно добавляем новые операции
- в одном месте делаем опрации над разными независимыми классами

Минусы:
- усложняет расширение иерархии классов, поскольку новые классы обычно требуют добавления нового метода visit для каждого посетителя.
*/

type shape interface {
	accept(visitor Visitor)
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) accept(visitor Visitor) {
	visitor.VisitForRectangle(r)
}

type Circle struct {
	Radius float64
}

func (c Circle) accept(visitor Visitor) {
	visitor.VisitForCircle(c)
}

type Visitor interface {
	VisitForRectangle(rect Rectangle)
	VisitForCircle(circle Circle)
}

type AreaCalculator struct {
	area float64
}

func (a AreaCalculator) VisitForRectangle(rect Rectangle) {
	a.area = rect.Width * rect.Height
	fmt.Println("rectangle area:", a.area)
}

func (a AreaCalculator) VisitForCircle(c Circle) {
	a.area = math.Pi * c.Radius * c.Radius
	fmt.Println("circle area:", a.area)
}

func main() {
	a := &AreaCalculator{}
	c := &Circle{
		Radius: 2,
	}
	c.accept(a)
	r := &Rectangle{
		Width:  2,
		Height: 2,
	}
	r.accept(a)
}
