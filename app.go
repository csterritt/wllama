package main

import (
	"context"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// PromptForResponse returns the result for running the given prompt
func (a *App) PromptForResponse(prompt string) string {
	chars := strings.Split(prompt, "")
	chLen := len(chars)
	newChars := make([]string, chLen)
	for i := 0; i < chLen; i += 1 {
		j := chLen - 1 - i
		newChars[i] = chars[j]
	}

	return strings.Join(newChars, "")
}
