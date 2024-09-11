//go:build js && wasm

package render

import (
	"syscall/js"
	"wasm/functions"
)

func RenderHome(this js.Value, args []js.Value) interface{} {
	functions.ConsoleLog("Render")
	content := getContent()

	content.Set("innerHTML", `
    <h1>Github Repository Search</h1>
    <div style="margin-bottom:16px;">
        <input id="search-input" type="text" placeholder="Enter search query" onkeyup="if(event.key === 'Enter') searchInput()">
        <button onclick="searchInput()">Search</button>
    </div>
    <div id="results"></div>
	`)

	return nil
}

func getContent() js.Value {
	document := js.Global().Get("document")

	return document.Call("getElementById", "content")

}
