package reflect

import "fmt"

func findType(b interface{}) {
	switch b.(type) {
	case int:
		fmt.Println("b type is int")
	case string:
		fmt.Println("b type is string")
	case float64:
		fmt.Println("b type is float64")
	case bool:
		fmt.Println("b type is bool")
	case Brand:
		fmt.Println("b type is Brand")
	case Event:
		fmt.Println("b type is Event")
	default:
		fmt.Println("Unknown Type")
	}
}

type Brand struct {
	BrandID   string
	BrandName string
}

type Event struct {
	EventID   int
	EventName string
}

type Market struct {
}

func ReflectDemo() {
	// var x int = 42
	// var y interface{} = 42

	var b interface{} = Brand{
		BrandID:   "beebet",
		BrandName: "Project Yankee",
	}
	findType(b)

	b = Event{
		EventID:   123,
		EventName: "Ind Vs Aus",
	}
	findType(b)

	b = Market{}
	findType(b)

	// t := reflect.TypeOf(x)
	// v := reflect.ValueOf(x)

	// fmt.Println("Type:", t)        // int
	// fmt.Println("Kind:", t.Kind()) // int
	// fmt.Println("Value:", v.Int()) // 42
}
