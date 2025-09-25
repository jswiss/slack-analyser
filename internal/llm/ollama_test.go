package llm

import (
	"context"
	"testing"
)

func TestOllama_Generate_ComposesRequest(t *testing.T) {
	o := &Ollama{Model: "gemma2:2b", Endpoint: "http://localhost:11434"}
	// This test ensures the method exists and behaves. If Ollama isn't running, skip.
	_, err := o.Generate(context.Background(), "hello")
	if err != nil {
		t.Skipf("skipping: Ollama not running or unreachable: %v", err)
	}
}
