//go:generate enumer -type=Level

package angels

// Level is an enumeration of levels that are allowed
// ENUM(
// Serafino, Cherubino, Troni
// Dominazione, Virtù, Potestà
// Principato, Arcangelo, Angelo
// )
type Level int

const (
	Serafino Level = iota
	Cherubino
	Troni
	Dominazione
	Virtù
	Potestà
	Principato
	Arcangelo
	Angelo
)
