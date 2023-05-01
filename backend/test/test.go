package test

type Test struct {
	Name     string
	Sections []Section
}

type Section struct {
	Name      string
	Exercises []any
}
