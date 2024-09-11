//go:build js && wasm

package main

import (
	"syscall/js"
	"wasm/functions"
)

var done chan struct{}

func main() {
	functions.ConsoleLog("WASM main function started")
	done = make(chan struct{})
	js.Global().Set("search", js.FuncOf(functions.Search))
	<-done
}
