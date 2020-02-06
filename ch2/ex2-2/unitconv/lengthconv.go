package uniconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gft", f) }
func (f Feet) String() string  { return fmt.Sprintf("%gm", m) }
