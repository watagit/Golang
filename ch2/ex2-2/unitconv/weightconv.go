package uniconv

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string    { return fmt.Sprintf("%gkg", k) }
func (k Kilogram) String() string { return fmt.Sprintf("%glb", p) }
