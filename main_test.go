package bonito

import (
	"fmt"
	"testing"
)

func TestAsSeparated(t *testing.T) {
	nums := map[int]string{
		1:       "1",
		11:      "11",
		121:     "121",
		1094:    "1,094",
		36875:   "36,875",
		109442:  "109,442",
		1890754: "1,890,754",
	}

	for k, v := range nums {
		s := Pretty(k).AsSeparated(",").str
		if s != v {
			t.Fatalf("expected %v, got %v", v, s)
		}
	}
}
