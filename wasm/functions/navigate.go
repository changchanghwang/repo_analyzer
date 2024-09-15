//go:build js && wasm

package functions

import "syscall/js"

func NavigateTo(path string) {
	window := js.Global().Get("window")
	history := window.Get("history")

	history.Call("pushState", nil, "", path)
	window.Call("renderPage")
}
