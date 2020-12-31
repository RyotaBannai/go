package other

type Person struct {
	name     string
	age      int
	Exported bool
}

var (
	Me = Person{name: "Ryota", age: 10, Exported: true}
)
