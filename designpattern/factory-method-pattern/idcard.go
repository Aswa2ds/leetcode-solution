package factory_method_pattern

import "fmt"

type IDCard struct {
	owner string
}

func (c *IDCard) use() {
	fmt.Println(c.owner + " uses IDCard!")
}

func (c *IDCard) getOwner() string {
	return c.owner
}

