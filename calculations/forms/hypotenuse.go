package forms
import "math"

// Erwartet die Längen der Katheten eines rechtwinkligen Dreiecks.
// Liefert die Länge der Hypotenuse.
func Hypotenuse(a, b float64) float64 {
	// TODO
	c := a*a+b*b
	return math.Sqrt(c)
}
