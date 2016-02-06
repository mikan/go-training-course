// Copyright 2015-2016 mikan. All rights reserved.

package lengthconv

import "fmt"

type Feat float64
type Meters float64

func (f Feat) String() string   { return fmt.Sprintf("%gft", f) }
func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
