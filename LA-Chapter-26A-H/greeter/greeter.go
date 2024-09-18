package greeter

type Greeter struct {
	Error   bool
	Message string
}

// func NewGreeter() *Greeter {
// 	return &Greeter{Message: "Hello"}
// }

func NewGreeter(name string) *Greeter {
	return &Greeter{Message: "Hello " + name}
}
