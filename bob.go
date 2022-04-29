package main

import (
	"log"
	"sync"
	"time"
)

type bob struct{}

func (b *bob) Name() string {
	return "Bob"
}

func (b *bob) Live() {
	b.eat()
	b.drink()
	b.poop()
	b.sleep()
}

func (b *bob) eat() {
	log.Println(b.Name(), "eating")
}

func (b *bob) drink() {
	log.Println(b.Name(), "drinking")
}

func (b *bob) poop() {
	log.Println(b.Name(), "pooping")
	// Mutex problem.
	m := &sync.Mutex{}
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
}

func (b *bob) sleep() {
	log.Println(b.Name(), "sleeping")
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(30 * time.Second)
		}()
	}
}
