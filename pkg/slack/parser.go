package slack

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ParseChannel reads a Slack export channel directory containing daily JSON files.
func ParseChannel(channelDir string) (Channel, error) {
	var channel Channel
	channel.Name = filepath.Base(channelDir)

	var dailyFiles []string
	err := filepath.WalkDir(channelDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(d.Name(), ".json") {
			dailyFiles = append(dailyFiles, path)
		}
		return nil
	})
	if err != nil {
		return Channel{}, err
	}

	sort.Strings(dailyFiles)
	for _, file := range dailyFiles {
		b, err := os.ReadFile(file)
		if err != nil {
			return Channel{}, fmt.Errorf("read %s: %w", file, err)
		}
		var raw []map[string]any
		if err := json.Unmarshal(b, &raw); err != nil {
			return Channel{}, fmt.Errorf("unmarshal %s: %w", file, err)
		}
		for _, m := range raw {
			if m["type"] != "message" {
				continue
			}
			tsStr, _ := m["ts"].(string)
			ts := parseSlackTS(tsStr)
			var threadPtr *time.Time
			if tts, ok := m["thread_ts"].(string); ok && tts != "" {
				t := parseSlackTS(tts)
				threadPtr = &t
			}
			user, _ := m["user"].(string)
			text, _ := m["text"].(string)
			channel.Messages = append(channel.Messages, Message{
				Type:     "message",
				UserID:   user,
				Text:     text,
				TS:       ts,
				ThreadTS: threadPtr,
			})
		}
	}

	return channel, nil
}

func parseSlackTS(ts string) time.Time {
	// Slack ts: seconds.nanoseconds as string
	if ts == "" {
		return time.Time{}
	}
	parts := strings.Split(ts, ".")
	sec, _ := strconv.ParseInt(parts[0], 10, 64)
	var nsec int64
	if len(parts) > 1 {
		// pad/truncate to nanoseconds
		frac := parts[1]
		if len(frac) > 9 {
			frac = frac[:9]
		} else if len(frac) < 9 {
			frac = frac + strings.Repeat("0", 9-len(frac))
		}
		n, _ := strconv.ParseInt(frac, 10, 64)
		nsec = n
	}
	return time.Unix(sec, nsec).UTC()
}

// GroupThreads clusters messages by thread root (thread_ts) with fallbacks for root messages.
func GroupThreads(messages []Message) []Thread {
	if len(messages) == 0 {
		return nil
	}
	// Map root ts (string) to thread
	index := make(map[string]*Thread)
	order := make([]string, 0)
	for _, m := range messages {
		var root time.Time
		if m.ThreadTS != nil && !m.ThreadTS.IsZero() {
			root = *m.ThreadTS
		} else {
			// message without thread_ts is its own root
			root = m.TS
		}
		key := root.Format(time.RFC3339Nano)
		th, ok := index[key]
		if !ok {
			index[key] = &Thread{RootTS: root}
			order = append(order, key)
			th = index[key]
		}
		th.Messages = append(th.Messages, m)
	}
	threads := make([]Thread, 0, len(order))
	for _, k := range order {
		th := index[k]
		sort.Slice(th.Messages, func(i, j int) bool { return th.Messages[i].TS.Before(th.Messages[j].TS) })
		threads = append(threads, *th)
	}
	sort.Slice(threads, func(i, j int) bool { return threads[i].RootTS.Before(threads[j].RootTS) })
	return threads
}
