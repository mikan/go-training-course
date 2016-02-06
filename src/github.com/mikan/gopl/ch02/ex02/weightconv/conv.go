// Copyright 2015-2016 mikan. All rights reserved.

package weightconv

// PToKG converts a Pounds length to Kilograms.
func PtoKG(m Pounds) Kilograms { return Kilograms(m * 0.45359237) }

// KGToP converts a Feat length to Kilograms.
func KGtoP(f Kilograms) Pounds { return Pounds(f * 2.2046) }
