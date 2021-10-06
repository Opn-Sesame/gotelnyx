package telnyx

import (
	"fmt"
	"io"
	"time"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

// RateLimitError is error for 429 http error
type RateLimitError struct {
	Reset time.Time
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("RateLimitError: reset at %v", e.Reset)
}
