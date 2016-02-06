// Copyright 2015-2016 mikan. All rights reserved.

package lengthconv

// MToF converts a Meters length to Feat.
func MtoF(m Meters) Feat { return Feat(m * 3.2808) }

// MToF converts a Feat length to Meters.
func FtoM(f Feat) Meters { return Meters(f * 0.3048) }
