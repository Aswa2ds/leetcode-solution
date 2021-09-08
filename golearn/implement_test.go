package golearn

import "testing"

type Duck interface {
	Walk()
}

type SmallDuck struct {
	Duck
}

func TestSmallDuck(t *testing.T) {
	sd := SmallDuck{}
	sd.Walk()
}
