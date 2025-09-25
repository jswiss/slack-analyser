package cli

import (
	"fmt"

	"github.com/jswiss/slack-analyser/internal/nl"
	an "github.com/jswiss/slack-analyser/pkg/analysis"
	slk "github.com/jswiss/slack-analyser/pkg/slack"
)

// Run loads a channel directory, plans from prompt, computes metric, and returns a printable string.
func Run(channelDir string, prompt string) (string, error) {
	ch, err := slk.ParseChannel(channelDir)
	if err != nil {
		return "", err
	}
	plan, err := nl.Plan(prompt)
	if err != nil {
		return "", err
	}
	threads := slk.GroupThreads(ch.Messages)
	switch plan.Metric {
	case nl.MetricAvgThreadLength:
		v := an.AverageThreadLength(threads)
		return fmt.Sprintf("Average thread length: %.2f", v), nil
	case nl.MetricAvgThreadDurationMinutes:
		v := an.AverageThreadDurationMinutes(threads)
		return fmt.Sprintf("Average thread duration (min): %.2f", v), nil
	case nl.MetricMostActiveUser:
		most, _ := an.MostAndLeastActiveUsers(ch.Messages)
		return fmt.Sprintf("Most active user: %s (%d messages)", most.UserID, most.Count), nil
	case nl.MetricLeastActiveUser:
		_, least := an.MostAndLeastActiveUsers(ch.Messages)
		return fmt.Sprintf("Least active user: %s (%d messages)", least.UserID, least.Count), nil
	case nl.MetricLongestMessage:
		lm := an.LongestMessage(ch.Messages)
		return fmt.Sprintf("Longest message by %s (%d chars)", lm.UserID, lm.Length), nil
	default:
		return "", fmt.Errorf("unsupported metric")
	}
}
