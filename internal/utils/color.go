package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToRGBInt(hex string) (int, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0, fmt.Errorf("invalid hex color: %s", hex)
	}
	value, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return 0, err
	}
	return int(value), nil
}
