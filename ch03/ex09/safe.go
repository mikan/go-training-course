package main

import "strconv"

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
