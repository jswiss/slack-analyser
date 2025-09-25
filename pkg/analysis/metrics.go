package analysis

import (
	"sort"
	"time"

	slk "github.com/jswiss/slack-analyser/pkg/slack"
)

func AverageThreadLength(threads []slk.Thread) float64 {
	if len(threads) == 0 {
		return 0
	}
	var total int
	for _, th := range threads {
		total += len(th.Messages)
	}
	return float64(total) / float64(len(threads))
}

func AverageThreadDurationMinutes(threads []slk.Thread) float64 {
	if len(threads) == 0 {
		return 0
	}
	var total time.Duration
	for _, th := range threads {
		if len(th.Messages) == 0 {
			continue
		}
		start := th.Messages[0].TS
		end := th.Messages[0].TS
		for _, m := range th.Messages {
			if m.TS.Before(start) {
				start = m.TS
			}
			if m.TS.After(end) {
				end = m.TS
			}
		}
		total += end.Sub(start)
	}
	return total.Minutes() / float64(len(threads))
}

type UserCount struct {
	UserID string
	Count  int
}

func MostAndLeastActiveUsers(messages []slk.Message) (UserCount, UserCount) {
	counts := map[string]int{}
	for _, m := range messages {
		if m.UserID == "" {
			continue
		}
		counts[m.UserID]++
	}
	list := make([]UserCount, 0, len(counts))
	for u, c := range counts {
		list = append(list, UserCount{UserID: u, Count: c})
	}
	if len(list) == 0 {
		return UserCount{}, UserCount{}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Count > list[j].Count })
	return list[0], list[len(list)-1]
}

type Longest struct {
	UserID string
	Text   string
	Length int
}

func LongestMessage(messages []slk.Message) Longest {
	var best Longest
	for _, m := range messages {
		l := len([]rune(m.Text))
		if l > best.Length {
			best = Longest{UserID: m.UserID, Text: m.Text, Length: l}
		}
	}
	return best
}
