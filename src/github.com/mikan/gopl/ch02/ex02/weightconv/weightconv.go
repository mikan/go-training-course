// Copyright 2015-2016 mikan. All rights reserved.

package weightconv

import "fmt"

type Pounds float64
type Kilograms float64

func (p Pounds) String() string     { return fmt.Sprintf("%glb", p) }
func (kg Kilograms) String() string { return fmt.Sprintf("%gkg", kg) }
