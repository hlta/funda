package utils

import (
	"fmt"
	"strconv"
)

const (
	Base10    = 10
	BitSize32 = 32
)

// ParseUint parses a string into a uint, assuming the string is a base 10 number with a maximum of 32 bits.
func ParseUint(value string) (uint, error) {
	parsed, err := strconv.ParseUint(value, Base10, BitSize32)
	if err != nil {
		return 0, fmt.Errorf("parse uint error: %w", err)
	}
	return uint(parsed), nil
}

func UintToString(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}
