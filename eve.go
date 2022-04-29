package main

import "log"

type eve struct {
	buffer [][MiB]byte
}

func (e *eve) Name() string {
	return "Eve"
}

func (e *eve) Live() {
	e.eat()
	e.drink()
	e.poop()
	e.sleep()
}

func (e *eve) eat() {
	log.Println(e.Name(), "eating")
	max := GiB
	for len(e.buffer)*MiB < max {
		e.buffer = append(e.buffer, [MiB]byte{})
	}
}

func (e *eve) drink() {
	log.Println(e.Name(), "drinking")
}

func (e *eve) poop() {
	log.Println(e.Name(), "poop")
}

func (e *eve) sleep() {
	log.Println(e.Name(), "sleeping")
}
