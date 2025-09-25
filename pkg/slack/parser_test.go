package slack_test

import (
	"path/filepath"
	"testing"

	slack "github.com/jswiss/slack-analyser/pkg/slack"
)

func TestParseChannel_BasicDailyFiles(t *testing.T) {
	// Arrange: use testdata export with two daily files
	exportRoot := filepath.Join("testdata", "exports", "channel_one")

	// Act
	channel, err := slack.ParseChannel(exportRoot)
	if err != nil {
		t.Fatalf("ParseChannel returned error: %v", err)
	}

	// Assert basic counts
	if channel.Name != "channel_one" {
		t.Fatalf("expected channel name 'channel_one', got %q", channel.Name)
	}

	if len(channel.Messages) != 6 {
		t.Fatalf("expected 6 messages aggregated across days, got %d", len(channel.Messages))
	}

	// Assert threading recognition: standalone messages are single-message threads → 3 threads
	threads := slack.GroupThreads(channel.Messages)
	if len(threads) != 3 {
		t.Fatalf("expected 3 threads, got %d", len(threads))
	}

	// There should be a thread of length 3, one of 2, and one of 1
	counts := []int{len(threads[0].Messages), len(threads[1].Messages), len(threads[2].Messages)}
	have := map[int]int{}
	for _, c := range counts {
		have[c]++
	}
	if have[3] != 1 || have[2] != 1 || have[1] != 1 {
		t.Fatalf("expected thread sizes 3,2,1; got %v", counts)
	}
}
