package main

import (
	"fmt"
	"os"

	"github.com/leodido/go-monorepo/amenadiel"
)

func main() {
	fmt.Fprintf(os.Stdout, "%s is an %s", amenadiel.Amenadiel.Name, amenadiel.Amenadiel.Level)
}
