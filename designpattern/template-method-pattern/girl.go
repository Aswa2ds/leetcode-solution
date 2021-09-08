package template_method_pattern

import "fmt"

type Girl struct {
	Person
}

func (_ *Girl) BeforeOut() {
	fmt.Println("dress up..")
}
