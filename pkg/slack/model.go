package slack

import "time"

type Message struct {
	Type     string
	UserID   string
	Text     string
	TS       time.Time
	ThreadTS *time.Time
}

type Channel struct {
	Name     string
	Messages []Message
}

type Thread struct {
	RootTS   time.Time
	Messages []Message
}
