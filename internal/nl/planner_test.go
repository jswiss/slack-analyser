package nl

import "testing"

func TestPlan_BaselineKeywords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  MetricKind
	}{
		{name: "avg thread length", input: "average thread length", want: MetricAvgThreadLength},
		{name: "avg thread duration", input: "average thread duration", want: MetricAvgThreadDurationMinutes},
		{name: "most active", input: "most active messager", want: MetricMostActiveUser},
		{name: "most chatty synonym", input: "who is the most chatty?", want: MetricMostActiveUser},
		{name: "least active", input: "least active user", want: MetricLeastActiveUser},
		{name: "longest message", input: "find the longest message", want: MetricLongestMessage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Plan(tt.input)
			if err != nil {
				t.Fatalf("Plan error: %v", err)
			}
			if got.Metric != tt.want {
				t.Fatalf("want %v, got %v", tt.want, got.Metric)
			}
		})
	}
}
