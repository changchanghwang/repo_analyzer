//go:build js && wasm

package functions

func evaluate(repoName string) {
	ConsoleLog("Showing details for:", repoName)
}
