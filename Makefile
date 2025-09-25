MODULE=github.com/jswiss/slack-analyser
OLLAMA?=ollama
OLLAMA_MODEL?=gemma2:2b

.PHONY: test race cover tidy deps ollama-start ollama-pull run demo

test:
	go test ./...

race:
	$(MAKE) ollama-start
	go test -race ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | tail -n 1

tidy:
	go mod tidy

deps:
	@echo "Ensure Ollama is installed and a Gemma model is pulled:"
	@echo "  ollama pull gemma2:2b"
	@echo "Or choose another Gemma variant."

ollama-start:
	@$(OLLAMA) serve >/dev/null 2>&1 &
	@sleep 1
	@echo "Ollama serve started (if not already running)"

ollama-pull:
	$(OLLAMA) pull $(OLLAMA_MODEL)

run:
	go run ./cmd/slack-analyser --channel-dir ./pkg/slack/testdata/exports/channel_one --prompt "average thread length"

demo:
	@echo "=== Slack Analyser Demo ==="
	@echo "1. Average thread length:"
	@go run ./cmd/slack-analyser --channel-dir ./pkg/slack/testdata/exports/channel_one --prompt "average thread length"
	@echo ""
	@echo "2. Most active user:"
	@go run ./cmd/slack-analyser --channel-dir ./pkg/slack/testdata/exports/channel_one --prompt "most active messager"
	@echo ""
	@echo "3. Longest message:"
	@go run ./cmd/slack-analyser --channel-dir ./pkg/slack/testdata/exports/channel_one --prompt "longest message"
	@echo ""
	@echo "4. Least active user:"
	@go run ./cmd/slack-analyser --channel-dir ./pkg/slack/testdata/exports/channel_one --prompt "least active user"


