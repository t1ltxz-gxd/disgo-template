// nolint:revive
package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseBytes takes a string with a byte size and returns the value to int64.
func ParseBytes(s string) (int64, error) {
	s = strings.TrimSpace(strings.ToLower(s))

	ordered := []struct {
		suf  string
		mult int64
	}{
		{"kb", 1 << 10},
		{"mb", 1 << 20},
		{"gb", 1 << 30},
		{"tb", 1 << 40},
		{"pb", 1 << 50},
		{"eb", 1 << 60},
		{"b", 1},
	}

	for _, o := range ordered {
		if strings.HasSuffix(s, o.suf) {
			numStr := strings.TrimSpace(strings.TrimSuffix(s, o.suf))
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return 0, fmt.Errorf("could not make out the number: %w", err)
			}
			return int64(num * float64(o.mult)), nil
		}
	}

	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("unknown size format: %w", err)
	}
	return n, nil
}
