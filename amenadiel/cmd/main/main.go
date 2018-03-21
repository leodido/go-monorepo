package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/leodido/go-monorepo/amenadiel"
	"github.com/leodido/go-monorepo/lucifer"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Fprintf(os.Stdout, "%s is an %s\n", amenadiel.Amenadiel.Name, amenadiel.Amenadiel.Level)
	fmt.Fprintf(os.Stdout, "He tries to take the %s back to its home - ie., the %s\n", lucifer.GetName(), lucifer.Home())
}
