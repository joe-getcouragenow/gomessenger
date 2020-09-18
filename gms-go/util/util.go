package util

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// UnixMilli ...
// many languages won't have nano or micro supported
// so a util for milliseconds will come in handy
func UnixMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

// NewTimedId ...
func NewTimedId() string {
	id := uuid.New().String()
	now := UnixMilli()
	return fmt.Sprintf("%v@%v", id, now)
}

// NewUserId ...
func NewUserId() string {
	id := NewTimedId()
	return fmt.Sprintf("%v$%v", id, "user")
}

// NewMessageId ...
func NewMessageId() string {
	id := NewTimedId()
	return fmt.Sprintf("%v$%v", id, "msg")
}
