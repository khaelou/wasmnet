package main

import (
	"fmt"
	"strconv"
	"syscall/js"

	"wasmnet/pool"
	"wasmnet/utils"
)

var (
	poolExecCount int
)

func registerCallbacks() {
	js.Global().Set("helloWorld", js.FuncOf(helloWorld))

	js.Global().Set("launchPool", js.FuncOf(launchPool))
	js.Global().Set("poolHandler", js.FuncOf(poolHandler))

	js.Global().Set("launchForLoop", js.FuncOf(launchForLoop))
	js.Global().Set("forHandler", js.FuncOf(forHandler))

	js.Global().Set("launchPeerJS", js.FuncOf(launchPeerJS))
	js.Global().Set("peerJSHandler", js.FuncOf(peerJSHandler))

	js.Global().Set("launchWebGL", js.FuncOf(launchWebGL))
}

func main() {
	c := make(chan struct{}, 0)

	println("Go WebAssembly")

	// Register js functions to go
	registerCallbacks()
	<-c
}

// Hello World example
func helloWorld(this js.Value, i []js.Value) interface{} {
	phrase := "Hello, world!"
	js.Global().Call("alert", phrase)

	return phrase
}

// Built-in task-runner / worker-pool
func launchPool(this js.Value, i []js.Value) interface{} {
	poolDiv := js.Global().Get("document").Call("getElementById", "poolDiv")
	js.Global().Get("document").Call("getElementById", "runPoolButton").Set("disabled", false)

	if poolDiv.Get("style").Get("display").String() == "none" {
		//fmt.Println("Show.1")
		poolDiv.Get("style").Set("display", "block")
	} else if poolDiv.Get("style").Get("display").String() == "block" {
		//fmt.Println("Hide.2")
		poolDiv.Get("style").Set("display", "none")
	} else {
		//fmt.Println("Show.0")
		poolDiv.Get("style").Set("display", "block")
	}

	return nil
}

// Handles input submitted via Frontend of form (forLoopDiv)
func poolHandler(this js.Value, i []js.Value) interface{} {
	//limitValue := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	//js.Global().Set("output", limitValue)

	outputTextArea := js.Global().Get("document").Call("getElementById", "poolOutput")
	outputTextArea.Set("value", "")

	poolExecCount += 1
	runCount := fmt.Sprintf("POOL_RUN %d", poolExecCount)

	utils.UpdatePoolOutput(runCount)
	pool.InitClient(100)

	return nil
}

// Show/Hide Frontend forLoopDiv
func launchForLoop(this js.Value, i []js.Value) interface{} {
	forLoopDiv := js.Global().Get("document").Call("getElementById", "forLoopDiv")
	js.Global().Get("document").Call("getElementById", "runForButton").Set("disabled", false)

	if forLoopDiv.Get("style").Get("display").String() == "none" {
		forLoopDiv.Get("style").Set("display", "block")
	} else if forLoopDiv.Get("style").Get("display").String() == "block" {
		forLoopDiv.Get("style").Set("display", "none")
	} else {
		forLoopDiv.Get("style").Set("display", "block")
	}

	return nil
}

// Handles input submitted via Frontend of form (forLoopDiv)
func forHandler(this js.Value, i []js.Value) interface{} {
	limitValue := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	js.Global().Set("output", limitValue)
	intLimit, _ := strconv.Atoi(limitValue)
	println("forHandler() limit:", intLimit)

	textAreaInput := js.Global().Get("document").Call("getElementById", "forInput")
	fmt.Println("TextArea Located:", textAreaInput.Get("name"))

	targetLine := fmt.Sprintf("%d", intLimit)
	utils.UpdateInputElement(targetLine) // update via Frontend

	if intLimit != 0 {
		for i := 0; i <= intLimit; i++ {
			targetLine := fmt.Sprintf("i = %d\n", i)
			fmt.Println(targetLine)

			utils.UpdateConsoleElement(targetLine) // update via Frontend
		}
	}

	return nil
}

// Show/Hide Frontend peerJSDiv
func launchPeerJS(this js.Value, i []js.Value) interface{} {
	peerJSDiv := js.Global().Get("document").Call("getElementById", "peerJSDiv")
	js.Global().Get("document").Call("getElementById", "runPeerJSButton").Set("disabled", false)

	if peerJSDiv.Get("style").Get("display").String() == "none" {
		peerJSDiv.Get("style").Set("display", "block")
	} else if peerJSDiv.Get("style").Get("display").String() == "block" {
		peerJSDiv.Get("style").Set("display", "none")
	} else {
		peerJSDiv.Get("style").Set("display", "block")
	}

	return nil
}

// Handles input submitted via Frontend of form (peerJSDiv)
func peerJSHandler(this js.Value, i []js.Value) interface{} {
	//targetURL := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()

	return nil
}

// Show/Hide Frontend webGLDiv
func launchWebGL(this js.Value, i []js.Value) interface{} {
	webGLDiv := js.Global().Get("document").Call("getElementById", "webGLDiv")

	if webGLDiv.Get("style").Get("display").String() == "none" {
		webGLDiv.Get("style").Set("display", "block")
	} else if webGLDiv.Get("style").Get("display").String() == "block" {
		webGLDiv.Get("style").Set("display", "none")
	} else {
		webGLDiv.Get("style").Set("display", "block")
	}

	// webGL.InitCanvas()

	return nil
}
