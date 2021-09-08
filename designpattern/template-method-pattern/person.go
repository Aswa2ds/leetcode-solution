package template_method_pattern

import "fmt"

type Person struct {
	Specific IPerson
	name string
}

func (this *Person) SetName(name string) {
	this.name = name
}

func (this *Person) Out() {
	this.BeforeOut()
	fmt.Println(this.name + " go out...")
}

func (this *Person) BeforeOut() {
	if this.Specific == nil {
		return
	}

	this.Specific.BeforeOut()
}
