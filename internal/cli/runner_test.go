package cli

import (
	"path/filepath"
	"testing"
)

func TestRun_Outputs(t *testing.T) {
	channelDir := filepath.Join("..", "..", "pkg", "slack", "testdata", "exports", "channel_one")

	tests := []struct {
		name         string
		prompt       string
		wantContains string
	}{
		{name: "avg length", prompt: "average thread length", wantContains: "2"},
		{name: "longest message", prompt: "longest message", wantContains: "U02"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Run(channelDir, tt.prompt)
			if err != nil {
				t.Fatalf("run error: %v", err)
			}
			if out == "" || !contains(out, tt.wantContains) {
				t.Fatalf("output %q does not contain %q", out, tt.wantContains)
			}
		})
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || (len(sub) > 0 && (indexOf(s, sub) >= 0)))
}

func indexOf(s, sub string) int {
	// simple implementation to avoid importing strings for the small test helper
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
