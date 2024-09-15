//go:build js && wasm

package render

import (
	"syscall/js"
	"wasm/functions"
)

func RenderEvaluate(this js.Value, args []js.Value) interface{} {
	functions.ConsoleLog("Render")
	content := getContent()

	content.Set("innerHTML", `
    <h1>Evaluate</h1>
    <div id="results"></div>
	`)

	return nil
}
