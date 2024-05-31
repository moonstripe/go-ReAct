package main

import (
	"fmt"
	"strings"
)

type Agent interface {
	SetSystemPrompt(query string)
	Run()
}

type ReActAgent struct {
	llmClient *OllamaClient
	tools     map[string]Tool
}

func NewReActAgent(llmClient *OllamaClient, tools []Tool) *ReActAgent {
	toolMap := make(map[string]Tool)
	for _, tool := range tools {
		toolMap[tool.Name()] = tool
	}
	return &ReActAgent{
		llmClient: llmClient,
		tools:     toolMap,
	}
}

func (agent *ReActAgent) ReasonAndAct(prompt string) (string, error) {
	// Query the LLM for reasoning
	reasoning, err := agent.llmClient.Query(prompt)
	if err != nil {
		return "", err
	}

	// Based on the reasoning, decide on an action
	actionPrompt := fmt.Sprintf("Based on the following reasoning, what should the agent do next?\nReasoning: %s\nAction:", strings.TrimSpace(reasoning))
	action, err := agent.llmClient.Query(actionPrompt)
	if err != nil {
		return "", err
	}

	action = strings.TrimSpace(action)

	// Check if the action involves using a tool
	for toolName, tool := range agent.tools {
		if strings.Contains(action, toolName) {
			// Extract input for the tool from the action
			input := strings.Replace(action, toolName, "", 1)
			input = strings.TrimSpace(input)
			toolResult, err := tool.Use(input)
			if err != nil {
				return "", err
			}
			return toolResult, nil
		}
	}

	return action, nil
}
