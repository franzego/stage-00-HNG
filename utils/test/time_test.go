package test

import (
	"testing"
	"time"

	"github.com/franzego/stage00/pkg"
)

func TestTimestampFormat(t *testing.T) {
	ts := pkg.Timestamp()
	if _, err := time.Parse(time.RFC3339, ts); err != nil {
		t.Fatalf("Timestamp returned invalid RFC3339: %v", err)
	}
}
