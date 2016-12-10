// Copyright 2016 mikan. All rights reserved.

package main

import "testing"

func TestOutline(t *testing.T) {
	err := outline("https://github.com/")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	// stdout is untestable.
}
