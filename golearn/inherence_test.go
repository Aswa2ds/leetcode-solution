package golearn

import (
	"fmt"
	"testing"
)

type Human struct {
	name string
	age int
}

type Girl struct {
	Human
}

type Boy struct {
	Human
}

func (h *Human) Speak() {
	fmt.Println("human")
}

func (b *Boy) Speak() {
	fmt.Println("Boy")
}

func (g *Girl) Speak() {
	fmt.Println("Girl")
}

func TestSpeak(t *testing.T) {

}