Slack Analyser (MVP)
====================

Analyze a Slack export with natural language. MVP supports single-channel analysis and computes:

- Average thread length
- Average thread duration
- Most active messager
- Least active messager
- Longest message

Tech stack
---------
- Golang
- Ollama + Gemma (local LLM)

Project layout
--------------
- `pkg/slack`: parsing Slack exports and grouping threads
- `pkg/analysis`: metrics over messages/threads
- `internal/nl`: NL planner (keyword baseline → metric)
- `internal/llm`: Ollama adapter for Gemma
- `cmd/slack-analyser`: CLI entrypoint (coming next)

Getting started
---------------
1) Requirements
- Go (matching `go.mod`)
- Ollama installed and running (`https://ollama.com`)
- Pull a Gemma model (example):

```
ollama pull gemma2:2b
```

2) Install dependencies

```
make tidy
```

3) Run tests

```
make test
```

Optional: race detector and coverage

```
make race
make cover
```

Usage (CLI)
-----------
CLI will accept a Slack channel export directory and a free-form prompt. Example (to be implemented in `cmd/slack-analyser`):

```
go run ./cmd/slack-analyser --channel-dir /path/to/channel \
  --prompt "average thread length"
```

Notes
-----
- User resolution (IDs → display names) is not implemented yet.
- Planner currently uses a deterministic keyword baseline; LLM integration is available via `internal/llm` for future prompts.


