package template_method_pattern

import "testing"

func TestBoy(t *testing.T) {
	var p = &Person{}

	p.Specific = &Boy{}
	p.SetName("Mike")
	p.Out()

}
