package class

type Human struct {
	name string
}

func NewHuman(name string) *Human {
	return &Human{name: name}
}
