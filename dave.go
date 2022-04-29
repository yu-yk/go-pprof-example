package main

import (
	"log"
	"time"
)

type dave struct{}

func (d *dave) Name() string {
	return "Dave"
}

func (d *dave) Live() {
	d.eat()
	d.drink()
	d.poop()
	d.sleep()
}

func (d *dave) eat() {
	log.Println(d.Name(), "eating")
}

func (d *dave) drink() {
	log.Println(d.Name(), "drinking")
	for i := 0; i < 10000000000; i++ {
		// Doing nothing.
	}
}

func (d *dave) poop() {
	log.Println(d.Name(), "pooping")
}

func (d *dave) sleep() {
	log.Println(d.Name(), "sleeping")
	// Blocks for one second
	<-time.After(time.Second)
}
