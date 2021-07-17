package main

func HandleData(x interface{}) {
	switch x.(type) {
	case int:
	case string:
	case map[string]int:
	}
}

func main() {
	HandleData(1)
	HandleData(map[string]int{"foo": 10})
	HandleData("Hello, Go!")
}