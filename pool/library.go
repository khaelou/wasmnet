package pool

import (
	"fmt"
	"math/rand"
	"time"
)

// Collection of built-in macros
var MacroLibrary = map[string]interface{}{ // Identifier: Function
	"HelloWorld": HelloWorld,
	"FooBar":     FooBar,
	"HeadsTails": HeadsTails,
}

func HelloWorld(i interface{}) interface{} {
	var output interface{}

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 5
	rand := rand.Intn(max-min+1) + min

	switch rand {
	case 1:
		output = "Hardware"
	case 2:
		output = "Software"
	case 3:
		output = "Home"
	case 4:
		output = "Auto"
	default:
		output = "Hello, world!"
	}

	if i != "" {
		return fmt.Sprintf("%s & %v", output, i)
	} else {
		return output
	}
}

func FooBar(i interface{}) interface{} {
	var output interface{}

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 3
	rand := rand.Intn(max-min+1) + min

	switch rand {
	case 1:
		output = "Foo"
	case 2:
		output = "Bar"
	default:
		output = "FooBar"
	}

	if i != "" {
		return fmt.Sprintf("%s & %v", output, i)
	} else {
		return output
	}
}

func HeadsTails(i interface{}) interface{} {
	var output interface{}

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 2
	rand := rand.Intn(max-min+1) + min

	switch rand {
	case 1:
		output = "Heads"
	case 2:
		output = "Tails"
	}

	if i != "" {
		return fmt.Sprintf("%s & %v", output, i)
	} else {
		return output
	}
}
