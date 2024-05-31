package main

import (
	"fmt"
)

type Tool interface {
	Name() string
	Use(input string) (string, error)
}

type ExampleTool struct{}

func (t *ExampleTool) Name() string {
	return "ExampleTool"
}

func (t *ExampleTool) Use(input string) (string, error) {
	// Example implementation: simply echoes the input.
	return fmt.Sprintf("ExampleTool used with input: %s", input), nil
}
