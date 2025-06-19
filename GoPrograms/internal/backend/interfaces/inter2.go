package backend

import "fmt"

func PrintAnything(v interface{}) {
	fmt.Println(v)
}

func TypeAssertion() {
	var s interface{} = "hello"
	str := s.(string) // Asserts that s is a string
	fmt.Println("String str is :", str)
	switch v := s.(type) {
	case string:
		fmt.Println("It's a string:", v)
	case int:
		fmt.Println("It's an int:", v)
	default:
		fmt.Println("Unknown type")
	}

}
