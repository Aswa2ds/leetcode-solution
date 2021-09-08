package template_method_pattern

import "fmt"

type Boy struct {
	Person
}

func (_ *Boy) BeforeOut() {
	fmt.Println("get up..")
}

