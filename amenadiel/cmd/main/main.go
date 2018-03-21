package main

import (
	"fmt"
	"os"

	"github.com/leodido/go-monorepo/amenadiel"
	"github.com/leodido/go-monorepo/lucifer"
)

func main() {
	fmt.Fprintf(os.Stdout, "%s is an %s", amenadiel.Amenadiel.Name, amenadiel.Amenadiel.Level)
	fmt.Fprintf(os.Stdout, "He tries to take the %s back to its home - ie., the %s", lucifer.GetName(), lucifer.Home())
}
