// elemPush.go: Updates frontend element values on-demand
package utils

import (
	"syscall/js"
)

// frontend utility function
func UpdatePoolOutput(addLine string) interface{} {
	js.Global().Call("addPoolOutput", addLine) // Call JavaScript function (output :=)
	//println("UPDATE_POOL_OUTPUT:", output.String())

	return nil
}

// called by: forHandler (frontend utility function)
func UpdateInputElement(addLine string) interface{} {
	add := js.Global().Call("addForInput", addLine) // Call JavaScript function
	println("Input TextArea:", add.String())

	consoleTextArea := js.Global().Get("document").Call("getElementById", "forOutput")
	consoleTextArea.Set("value", "")

	//js.Global().Get("config").Set("key", "12345")

	return nil
}

// called by: forHandler (frontend utility function)
func UpdateConsoleElement(addLine string) interface{} {
	add := js.Global().Call("addForConsoleOutput", addLine) // Call JavaScript function
	println("Console TextArea:", add.String())

	return nil
}
