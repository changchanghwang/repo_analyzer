//go:build js && wasm

package main

import (
	"syscall/js"
	"wasm/functions"
	"wasm/render"
)

var done chan struct{}

func main() {
	functions.ConsoleLog("WASM main function started")
	done = make(chan struct{})
	js.Global().Set("renderHome", js.FuncOf(render.RenderHome))
	js.Global().Set("renderEvaluate", js.FuncOf(render.RenderEvaluate))
	js.Global().Set("search", js.FuncOf(functions.Search))
	<-done
}
