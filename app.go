package main

import (
	"context"
	"io"
	"os/exec"
)

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
	_, err = stdin.Write([]byte(prompt + "\n"))
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

	return string(data)
}
