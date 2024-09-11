//go:build js && wasm

package functions

import "syscall/js"

func evaluate(repoName string) {
	window := js.Global().Get("window")
	window.Get("location").Set("href", "/evaluate?repoName="+repoName)
}
