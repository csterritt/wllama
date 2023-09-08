package main

import (
	"context"
	"io"
	"os/exec"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

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

	// The policy can then be used to sanitize lots of input and it is safe to use the policy in multiple goroutines
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
func (a *App) PromptForResponse(prompt string) string {
	cmd := exec.Command("/opt/homebrew/bin/ollama", "run", "codellama:7b-instruct")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "cmd setup failed with error " + err.Error()
	}

	stdin, _ := cmd.StdinPipe()
	oneLinePrompt := strings.Join(strings.Split(prompt, "\n"), " ")
	_, err = stdin.Write([]byte(oneLinePrompt + "\n"))
	if err != nil {
		return "stdin.Write failed with error " + err.Error()
	}

	err = stdin.Close()
	if err != nil {
		return "stdin.Close failed with error " + err.Error()
	}

	err = cmd.Start()
	if err != nil {
		return "cmd.Start failed with error " + err.Error()
	}

	data, err := io.ReadAll(stdout)
	if err != nil {
		return "ReadAll failed with error " + err.Error()
	}

	if err := cmd.Wait(); err != nil {
		return "cmd Wait failed with error " + err.Error()
	}

	return convertToHtml(string(data))
}
