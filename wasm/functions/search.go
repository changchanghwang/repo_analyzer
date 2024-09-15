//go:build js && wasm

package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"syscall/js"
)

func ConsoleLog(args ...interface{}) {
	js.Global().Get("console").Call("log", args...)
}

type searchResult struct {
	TotalCount int `json:"total_count"`
	Items      []struct {
		FullName        string `json:"full_name"`
		Description     string `json:"description"`
		HtmlURL         string `json:"html_url"`
		ForksCount      int    `json:"forks_count"`
		OpenIssuesCount int    `json:"open_issues_count"`
		StargazersCount int    `json:"stargazers_count"`
	} `json:"items"`
}

func Search(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	input := document.Call("getElementById", "search-input")
	value := input.Get("value").String()

	go func() {
		resp, err := http.Get("http://localhost:3334/api/search?q=" + value)
		if err != nil {
			ConsoleLog("Error making HTTP request:", err.Error())
			displayError(document, err.Error())
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			ConsoleLog("Error reading response body:", err.Error())
			displayError(document, err.Error())
			return
		}

		var result searchResult
		err = json.Unmarshal(body, &result)
		if err != nil {
			ConsoleLog("Error unmarshaling JSON:", err.Error())
			displayError(document, err.Error())
			return
		}

		displayResults(document, &result)
	}()

	return nil
}

func displayError(document js.Value, message string) {
	ConsoleLog("Displaying error:", message)
	resultDiv := document.Call("getElementById", "results")
	resultDiv.Set("innerHTML", "Error: "+message)
}

func displayResults(document js.Value, result *searchResult) {
	resultDiv := document.Call("getElementById", "results")
	resultDiv.Set("innerHTML", "") // Clear previous results

	table := document.Call("createElement", "table")
	table.Get("style").Set("width", "100%")
	table.Get("style").Set("borderCollapse", "collapse")

	// Create table header
	thead := document.Call("createElement", "thead")
	headerRow := document.Call("createElement", "tr")
	headers := []string{"Repository", "Description", "Stars", "Forks", "Open Issues", "Analyze"}

	for _, header := range headers {
		th := document.Call("createElement", "th")
		th.Set("innerHTML", header)
		th.Get("style").Set("border", "1px solid black")
		th.Get("style").Set("padding", "8px")
		headerRow.Call("appendChild", th)
	}
	thead.Call("appendChild", headerRow)
	table.Call("appendChild", thead)

	// Create table body
	tbody := document.Call("createElement", "tbody")
	for _, item := range result.Items {
		row := document.Call("createElement", "tr")

		createCell := func(content string) {
			td := document.Call("createElement", "td")
			td.Set("innerHTML", content)
			td.Get("style").Set("border", "1px solid black")
			td.Get("style").Set("padding", "8px")
			row.Call("appendChild", td)
		}

		createCell(fmt.Sprintf("<a href='%s' target='_blank'>%s</a>", item.HtmlURL, item.FullName))
		createCell(item.Description)
		createCell(fmt.Sprintf("%d", item.StargazersCount))
		createCell(fmt.Sprintf("%d", item.ForksCount))
		createCell(fmt.Sprintf("%d", item.OpenIssuesCount))

		// Add button
		buttonCell := document.Call("createElement", "td")
		buttonCell.Get("style").Set("border", "1px solid black")
		buttonCell.Get("style").Set("text-align", "center")
		button := document.Call("createElement", "button")
		button.Set("innerHTML", "Analyze")
		button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			NavigateTo("/evaluate?repoName=" + item.FullName)
			return nil
		}))
		buttonCell.Call("appendChild", button)
		row.Call("appendChild", buttonCell)

		tbody.Call("appendChild", row)
	}
	table.Call("appendChild", tbody)

	resultDiv.Call("appendChild", table)
}
