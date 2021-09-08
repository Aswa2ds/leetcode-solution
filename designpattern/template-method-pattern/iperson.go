package template_method_pattern

type IPerson interface {
	SetName(name string)
	BeforeOut()
	Out()
}
