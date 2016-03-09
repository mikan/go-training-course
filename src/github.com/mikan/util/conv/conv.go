// Copyright 2016 mikan. All rights reserved.

package conv

import "strconv"

// Converts string value, or returns default if value is empty or not convertible.
func SafeAtoI(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	ss, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return ss
}

func SafeRangedAtoI(s string, defaultValue, min, max int) int {
	if s == "" {
		return defaultValue
	}
	ss, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	if ss < min {
		return min
	}
	if ss > max {
		return max
	}
	return ss
}

// Converts string value, or returns default if value is empty or not convertible.
func SafeAtoF(s string, defaultValue float64) float64 {
	if s == "" {
		return defaultValue
	}
	ss, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultValue
	}
	return ss
}
