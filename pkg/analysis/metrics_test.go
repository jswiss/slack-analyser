package analysis_test

import (
	"path/filepath"
	"testing"

	an "github.com/jswiss/slack-analyser/pkg/analysis"
	slk "github.com/jswiss/slack-analyser/pkg/slack"
)

func loadMessages(t *testing.T) []slk.Message {
	t.Helper()
	root := filepath.Join("..", "slack", "testdata", "exports", "channel_one")
	ch, err := slk.ParseChannel(root)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	return ch.Messages
}

func TestAverageThreadLength(t *testing.T) {
	msgs := loadMessages(t)
	threads := slk.GroupThreads(msgs)
	got := an.AverageThreadLength(threads) // sizes: 3,2,1 → avg 2.0
	if got != 2 {
		t.Fatalf("expected 2.0, got %v", got)
	}
}

func TestAverageThreadDurationMinutes(t *testing.T) {
	msgs := loadMessages(t)
	threads := slk.GroupThreads(msgs)
	got := an.AverageThreadDurationMinutes(threads)
	if got <= 0 || got >= 120 {
		t.Fatalf("unexpected avg duration: %v", got)
	}
}

func TestMostAndLeastActiveMessager(t *testing.T) {
	msgs := loadMessages(t)
	most, least := an.MostAndLeastActiveUsers(msgs)
	if most.UserID == "" || least.UserID == "" {
		t.Fatalf("expected non-empty user IDs")
	}
	if most.Count < least.Count {
		t.Fatalf("most should have >= least: most=%v least=%v", most, least)
	}
}

func TestLongestMessage(t *testing.T) {
	msgs := loadMessages(t)
	lm := an.LongestMessage(msgs)
	if lm.UserID != "U02" {
		t.Fatalf("expected U02 to have longest message, got %s", lm.UserID)
	}
}
