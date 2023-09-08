package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

type RequestToOllama struct {
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
}

type ResponsePart struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Done      bool   `json:"done"`
	Response  string `json:"response"`
	Context   []int  `json:"context"`
}

var policy *bluemonday.Policy

func init() {
	policy = bluemonday.UGCPolicy()
}

func convertToHtml(output string) string {
	allLines := strings.Split(output, "\n")
	lines := make([]string, 0)
	for _, s := range allLines {
		t := strings.Trim(s, " ")
		if len(t) > 0 {
			lines = append(lines, t)
		}
	}

	// The policy can then be used to sanitize lots of input, and it is safe to use the policy in multiple goroutines
	result := "<div>" + strings.Join(lines, "</div><div>") + "</div>"
	result = policy.Sanitize(
		result,
	)

	return result
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// PromptForResponse returns the result for running the given prompt
func (a *App) PromptForResponse(model string, prompt string) string {
	res, err := doPromptForResponse(model, prompt)
	if err != nil {
		return "Error caught: " + err.Error()
	}

	return res
}

func doPromptForResponse(model string, prompt string) (string, error) {
	url := "http://localhost:11434/api/generate"
	args := RequestToOllama{
		Prompt: prompt,
		Model:  model,
	}

	jsonData, err := json.Marshal(args)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/ndjson; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer func() {
		err = errors.Join(err, response.Body.Close()) // Magic!
	}()

	reader := bufio.NewReader(response.Body)
	scanner := bufio.NewScanner(reader)
	result := ""
	var lineResponse ResponsePart
	for scanner.Scan() {
		line := scanner.Text()

		err = json.Unmarshal([]byte(line), &lineResponse)
		if err != nil {
			return "", err
		}

		if !lineResponse.Done {
			result = result + lineResponse.Response
		}
	}

	if err = scanner.Err(); err != nil {
		return "", err
	}

	return convertToHtml(result), nil
}
