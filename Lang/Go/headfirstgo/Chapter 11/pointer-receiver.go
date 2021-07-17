package main

type Hello interface {
	SayHello()
}

type HelloGo string

func (h *HelloGo) SayHello() {}

func main() {
	helloGo := HelloGo("Hello, world")
	var hello Hello = &helloGo

	hello.SayHello()
}
