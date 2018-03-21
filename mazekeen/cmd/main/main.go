package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/leodido/go-monorepo/lucifer"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Fprintf(os.Stdout, "Mazekeen works for %s\n", lucifer.GetName())
}
