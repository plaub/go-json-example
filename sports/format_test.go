package sports_test

import (
	"testing"

	"github.com/plaub/go-json-example/sports"
)

func TestSecondsToTimeString(t *testing.T) {
	tests := []struct {
		seconds int64
		want    string
	}{
		{0, "00:00:00"},
		{1, "00:00:01"},
		{60, "00:01:00"},
		{61, "00:01:01"},
		{3599, "00:59:59"},
		{3600, "01:00:00"},
		{3601, "01:00:01"},
		{3661, "01:01:01"},
		{86399, "23:59:59"},
		{86400, "24:00:00"},
		{86401, "24:00:01"},
		{86461, "24:01:01"},
		{90061, "25:01:01"},
	}

	for _, test := range tests {
		got := sports.SecondsToTimeString(test.seconds)
		if got != test.want {
			t.Errorf("Expected %s, got %s", test.want, got)
		}
	}
}
