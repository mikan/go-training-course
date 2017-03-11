// Copyright 2015-2016 mikan. All rights reserved.

package main

import (
	"bytes"
)

func ExampleMain() {
	main()
	// Output:
	// Input characters (Ctrl-D to end) >>>
	// rune	count
	//
	// len	count
	// 1	0
	// 2	0
	// 3	0
	// 4	0
}

func ExampleCountWith1Len() {
	count(bytes.NewBufferString("A"))
	// Output:
	// Input characters (Ctrl-D to end) >>>
	// rune	count
	// 'A'	1
	//
	// len	count
	// 1	1
	// 2	0
	// 3	0
	// 4	0
}

func ExampleCountWith2Len() {
	count(bytes.NewBufferString("éé"))
	// Output:
	// Input characters (Ctrl-D to end) >>>
	// rune	count
	// 'é'	2
	//
	// len	count
	// 1	0
	// 2	2
	// 3	0
	// 4	0
}

func ExampleCountWith3Len() {
	count(bytes.NewBufferString("あああ"))
	// Output:
	// Input characters (Ctrl-D to end) >>>
	// rune	count
	// 'あ'	3
	//
	// len	count
	// 1	0
	// 2	0
	// 3	3
	// 4	0
}
