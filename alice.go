package main

import "log"

type alice struct {
}

func (a *alice) Name() string {
	return "Alice"
}

func (a *alice) Live() {
	a.eat()
	a.drink()
	a.poop()
	a.sleep()
}

func (a *alice) eat() {
	log.Println(a.Name(), "eating")
}

func (a *alice) drink() {
	log.Println(a.Name(), "drinking")
}

func (a *alice) poop() {
	log.Println(a.Name(), "pooping")
	// Produces a lot of garbage.
	_ = make([]byte, 16*MiB)
}

func (a *alice) sleep() {
	log.Println(a.Name(), "sleeping")
}
