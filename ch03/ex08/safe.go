package main

import "strconv"

// SafeAtoI converts string value, or returns default if value is empty or not convertible.
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

// SafeAtoF converts string value, or returns default if value is empty or not convertible.
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

// SafeRangedAtoI converts string value, or returns min if value is smaller than
// min, returns max if value is larger than max, or returns defaultValue if value
// is empty, not convertible.
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
