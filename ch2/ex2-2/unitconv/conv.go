package uniconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func PToK(p Pound) Kilogram { return Kilogram(p / 2.20462) }
func KToP(k Kilogram) Pound { return Pound(k * 2.20462) }

func FToM(f Feet) Meter { return Meter(f * 0.3048) }
func MToF(m Meter) Feet { return Feet(m / 0.3048) }
