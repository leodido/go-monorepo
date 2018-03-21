package lucifer

import (
	"math/rand"

	"github.com/leodido/go-monorepo/amenadiel"
	"github.com/leodido/go-monorepo/angels"
)

var names = [...]string{
	"Lucifero",
	"Belzebu",
	"Satana",
	"Bestia",
	"Mefistofele",
	"Demonio",
	"Diavolo",
}

// GetName ...
func GetName() string {
	return names[rand.Intn(len(names)-0)+0]
}

// Fights ....
func Fights() *angels.Angel {
	return amenadiel.Amenadiel
}

// Home ...
func Home() string {
	return "hell"
}
