package cli

import (
	"path/filepath"
	"testing"
)

func TestIntegration_EndToEnd(t *testing.T) {
	channelDir := filepath.Join("..", "..", "pkg", "slack", "testdata", "exports", "channel_one")

	tests := []struct {
		name         string
		prompt       string
		wantContains []string
	}{
		{
			name:         "average thread length",
			prompt:       "average thread length",
			wantContains: []string{"Average thread length", "2.00"},
		},
		{
			name:         "most active user",
			prompt:       "most active messager",
			wantContains: []string{"Most active user", "messages"},
		},
		{
			name:         "longest message",
			prompt:       "longest message",
			wantContains: []string{"Longest message by", "U02"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Run(channelDir, tt.prompt)
			if err != nil {
				t.Fatalf("Run error: %v", err)
			}
			if out == "" {
				t.Fatalf("expected non-empty output")
			}
			for _, want := range tt.wantContains {
				if !contains(out, want) {
					t.Errorf("output %q does not contain %q", out, want)
				}
			}
		})
	}
}
