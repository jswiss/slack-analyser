package nl

import (
	"errors"
	"strings"
)

type MetricKind int

const (
	MetricUnknown MetricKind = iota
	MetricAvgThreadLength
	MetricAvgThreadDurationMinutes
	MetricMostActiveUser
	MetricLeastActiveUser
	MetricLongestMessage
)

type PlanRequest struct {
	Metric MetricKind
}

// Plan implements a minimal keyword baseline mapping NL prompts to metrics.
func Plan(input string) (PlanRequest, error) {
	q := strings.ToLower(strings.TrimSpace(input))
	switch {
	case containsAny(q, []string{"avg thread length", "average thread length"}):
		return PlanRequest{Metric: MetricAvgThreadLength}, nil
	case containsAny(q, []string{"avg thread duration", "average thread duration", "thread duration"}):
		return PlanRequest{Metric: MetricAvgThreadDurationMinutes}, nil
	case containsAny(q, []string{"most active", "most messages", "top poster", "most chatty"}):
		return PlanRequest{Metric: MetricMostActiveUser}, nil
	case containsAny(q, []string{"least active", "fewest messages", "least chatty"}):
		return PlanRequest{Metric: MetricLeastActiveUser}, nil
	case containsAny(q, []string{"longest message", "longest msg"}):
		return PlanRequest{Metric: MetricLongestMessage}, nil
	}
	return PlanRequest{}, errors.New("could not plan metric from input")
}

func containsAny(s string, needles []string) bool {
	for _, n := range needles {
		if strings.Contains(s, n) {
			return true
		}
	}
	return false
}
