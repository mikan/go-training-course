// Copyright 2016 mikan. All rights reserved.

package main

import "testing"

func TestMakeSHA256Length(t *testing.T) {
	actual := makeSHA("foo", 256)
	b := []byte(actual)
	if len(b)*8 != 256 {
		t.Errorf("SHA256 length isn't 256: %v (l=%d)", actual, len(b)*8)
	}
}

func TestMakeSHA384Length(t *testing.T) {
	actual := makeSHA("foo", 384)
	b := []byte(actual)
	if len(b)*8 != 384 {
		t.Errorf("SHA384 length isn't 384: %v (l=%d)", actual, len(b)*8)
	}
}

func TestMakeSHA512Length(t *testing.T) {
	actual := makeSHA("foo", 512)
	b := []byte(actual)
	if len(b)*8 != 512 {
		t.Errorf("SHA512 length isn't 512: %v (l=%d)", actual, len(b)*8)
	}
}
