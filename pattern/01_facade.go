package main

import (
	"fmt"
	"time"
)

/*
Суть паттерна "фасад" — реализовать простой доступ к сложной системе,
скрыть лишнее и избавить пользователя от необходимости глубоко разбираться
в устройстве системы.

Плюсы:
- Сокрытие сложной логики и простой интерфейс

Минусы:
- Интерфейс станет супер-объектом, к которому всё будет привязано

Реализуем мини систему домашнего кинотеатра. Наш фасад запускает просмотр фильма и заканчивает,
скрывая всю реализацию от клиента внутри себя
*/

type DVDPlayer struct {
}

func (dvd *DVDPlayer) On() {
	time.Sleep(time.Second)
	fmt.Println("DVDPlayer On")
}

func (dvd *DVDPlayer) Off() {
	time.Sleep(time.Second)
	fmt.Println("DVDPlayer Off")
}

type Projector struct{}

func (projector *Projector) On() {
	time.Sleep(time.Second)
	fmt.Println("Projector On")
}
func (projector *Projector) Off() {
	time.Sleep(time.Second)
	fmt.Println("Projector Off")
}

type SoundSystem struct{}

func (soundSystem *SoundSystem) On() {
	time.Sleep(time.Second)
	fmt.Println("SoundSystem On")
}

func (soundSystem *SoundSystem) Off() {
	time.Sleep(time.Second)
	fmt.Println("SoundSystem Off")
}

type HomeTheaterFacade struct {
	DvdPlayer   *DVDPlayer
	Projector   *Projector
	SoundSystem *SoundSystem
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{DvdPlayer: &DVDPlayer{},
		Projector:   &Projector{},
		SoundSystem: &SoundSystem{}}
}

func (htf *HomeTheaterFacade) WatchFilm() {
	htf.DvdPlayer.On()
	htf.Projector.On()
	htf.SoundSystem.On()
	fmt.Println("Starting watch film")
	time.Sleep(2 * time.Second)
}

func (htf *HomeTheaterFacade) EndFilm() {
	fmt.Println("Ending watch film")
	htf.DvdPlayer.Off()
	htf.Projector.Off()
	htf.SoundSystem.Off()
}

func main() {
	theater := NewHomeTheaterFacade()
	theater.WatchFilm()
	fmt.Println()
	theater.EndFilm()
}
